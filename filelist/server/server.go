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
			log.Panic(errors.New("not implemented"))
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
	fmt.Printf("saving: %s\n", targetTorrentFile)

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