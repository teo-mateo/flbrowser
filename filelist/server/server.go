package server

import (
	"net/http"
	"errors"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/teo-mateo/flbrowser/filelist/browse"
	"github.com/gorilla/handlers"
	"log"
	"github.com/teo-mateo/flbrowser/filelist/rtorrent"
	"strconv"
	"github.com/teo-mateo/flbrowser/filelist"
	"path"
	"os"
	"io/ioutil"
	"encoding/json"
	"sort"
	"time"
)

func httpError (err error, w http.ResponseWriter){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

//do not set here
var apiKey string = ""

/*
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
*/

func Start(port int, key string, username string, pwd string, clientDir string){

	apiKey = key
	if apiKey == ""{
		panic(errors.New("API Key missing"))
	}

	router := mux.NewRouter()

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){

		if checkAccessToken(r.Header.Get(ATHeader)){
			//redirect to /
			w.WriteHeader(http.StatusMovedPermanently)
			w.Header().Set("Location", "/")
		} else {
			err := doBasicAuth(w, r, username, pwd)
			if err != nil {
				fmt.Println(err.Error())
				w.Header().Set("WWW-Authenticate", "Basic")
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				//redirect to /
				w.WriteHeader(http.StatusMovedPermanently)
				w.Header().Set("Location", "/")
			}
		}

	}).Methods("GET")

	router.HandleFunc("/login2", func(w http.ResponseWriter, r *http.Request){
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil{
			httpError(err, w)
			return
		}
		var login map[string]string
		json.Unmarshal(b, &login)

		var u, p string
		var ok bool
		if u, ok = login["username"]; !ok{
			httpError(errors.New("username is mandatory"), w)
			return
		}
		if p, ok = login["password"]; !ok{
			httpError(errors.New("password is mandatory"), w)
			return
		}

		if u == username && p == pwd {
			//generate access token
			at, expires, err := generateAccessToken()
			if err != nil{
				httpError(err, w)
				return
			}
			//generate ad-hoc struct to send back access token
			response := struct{
				AccessToken string `json:"accessToken"`
				Expires time.Time `json:"expires"`
			} {
				at, expires,
			}

			//marshal and reply with the access token
			b, err := json.Marshal(response)
			w.Write(b)
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

	}).Methods("POST")

	router.HandleFunc("/ping",secure(ping)).Methods("GET")
	router.HandleFunc("/categories", secure(getFLCategories)).Methods("GET", "OPTIONS")
	router.HandleFunc("/torrents/fl/{category}/{page}", secure(listFLTorrents)).Methods("GET")
	router.HandleFunc("/torrents/rtr", secure(listRTRTorrents)).Methods("GET")
	router.HandleFunc("/torrents/fl/{id}/download", downloadTorrent).Methods("POST")
	router.HandleFunc("/torrents/rtr/{id}/{action}", secure(doRTRAction)).Methods("POST")


	//serve static files
	router.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir(clientDir))))

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

	//allow CORS

	opts := []handlers.CORSOption {

		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "FLAccessToken"}),
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(opts...)(router)))

}

func ping (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("pong"))
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

func getFLCategories(w http.ResponseWriter, r *http.Request){
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
