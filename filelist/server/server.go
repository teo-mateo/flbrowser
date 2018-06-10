package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/teo-mateo/flbrowser/filelist/browse"
	"github.com/teo-mateo/flbrowser/filelist"
	"github.com/teo-mateo/flbrowser/filelist/rtorrent"
	"log"
	"path"
	"os"
	"io/ioutil"
	"sort"
)

func httpError (err error, w http.ResponseWriter){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

//do not set here
var apiKey string = ""

func checkKey(w http.ResponseWriter, r *http.Request) bool {
	if apiKey == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	if r.Header.Get("APIKEY") != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	} else {
		return true
	}
}

func Start(port int, key string){

	apiKey = key
	if apiKey == ""{
		panic(errors.New("API Key missing"))
	}

	router := mux.NewRouter()

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request){

		categories := make([]browse.Category, 0)
		for _, v := range browse.Categories{
			categories = append(categories, v)
		}

		//sort categories, order by ID
		sort.Slice(categories, func(i int,j int) bool {
			return categories[i].ID < categories[j].ID
		})

		bytes, err := json.MarshalIndent(categories, " ", " ")
		if err != nil{
			httpError(err, w)
			return
		}
		_, err = w.Write(bytes)
		if err != nil{
			httpError(err, w)
			return
		}
	}).Methods("GET")

	router.HandleFunc("/torrents/fl/{category}/{page}", func(w http.ResponseWriter, r *http.Request) {
		if checkKey(w, r){
			listFLTorrents(w, r)
		}}).Methods("GET")

	router.HandleFunc("/torrents/rtr", func (w http.ResponseWriter, r *http.Request){
		if checkKey(w, r) {
			listRTRTorrents(w, r)
		}}).Methods("GET")

	router.HandleFunc("/torrents/fl/{id}/download", func (w http.ResponseWriter, r *http.Request){
		if checkKey(w, r){
			downloadTorrent(w, r)
		}}).Methods("POST")

	router.HandleFunc("/torrents/rtr/{id}/{action}", func (w http.ResponseWriter, r *http.Request){
		if checkKey(w, r){
			doRTRAction(w, r)
		}}).Methods("POST")

	fmt.Printf("Listening @ 127.0.0.1:%d\n", port)
	fmt.Println("Routes:")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil{
			return err
		}
		fmt.Printf("  %s\n", path)
		return nil
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}

func doRTRAction(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		httpError(errors.New("missing rtr id"), w)
		return
	}

	action := vars["action"]
	if action == ""{
		httpError(errors.New("mission rtr action"), w)
		return
	}

	rtrFunction := ""
	switch action {
	case "close":
		rtrFunction = "d.close"
		break
	case "open":
		rtrFunction = "d.open"
		break
	case "resume":
		rtrFunction = "d.resume"
		break
	case "pause":
		rtrFunction = "d.pause"
		break
	case "start":
		rtrFunction = "d.start"
		break
	case "stop":
		rtrFunction = "d.stop"
		break
	case "erase":
		rtrFunction = "d.erase"
		break
	default:
		httpError(errors.New("unknown action"), w)
		return
	}

	if rtrFunction != ""{
		_, err := rtorrent.RPC_id__bool(rtrFunction, id)
		if err != nil{
			httpError(err, w)
			return
		}
	}
}

func downloadTorrent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		httpError(err, w)
		return
	}

	filename, bytes, err := filelist.DownloadTorrent(id)
	if err != nil{
		httpError(err, w)
		return
	}

	//log some info
	fmt.Printf("Downloaded torrent file %s, bytes: %d\n", filename, len(bytes))

	var targetTorrentFile = path.Join(rtorrent.RActive, filename)
	fmt.Printf("Saving torrent: %s\n", targetTorrentFile)

	if _, err := os.Stat(targetTorrentFile); os.IsNotExist(err){
		err = ioutil.WriteFile(targetTorrentFile, bytes, 0644)
		if err != nil{
			httpError(err, w)
			return
		}
	} else {
		fmt.Println("...torrent exists")
	}
}

func listFLTorrents(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	category, err := strconv.Atoi(vars["category"])
	if err != nil{
		httpError(err, w)
		return
	}
	page, err := strconv.Atoi(vars["page"])
	if err != nil{
		httpError(err, w)
		return
	}

	if !browse.IsCategory(category){
		httpError(errors.New(fmt.Sprintf("%d is not a category", category)), w)
		return
	}

	torrents, err := filelist.GetTorrents(category, page)
	if err != nil {
		httpError(err, w)
		return
	}

	bytes, err := json.MarshalIndent(torrents, "", "  ")
	if err != nil {
		httpError(err, w)
		return
	}

	w.Write(bytes)
}

func listRTRTorrents(w http.ResponseWriter, r *http.Request){
	torrents, err := rtorrent.GetTorrents()
	if err != nil{
		httpError(err, w)
		return
	}

	bytes, err := json.MarshalIndent(torrents, "", "  ")
	if err != nil{
		httpError(err, w)
		return
	}

	w.Write(bytes)
}