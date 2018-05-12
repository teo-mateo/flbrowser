package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"errors"
	"fmt"
	"log"
	"strconv"
	"encoding/json"
	"github.com/teo-mateo/flbrowser/filelist/browse"
	"github.com/teo-mateo/flbrowser/filelist"
)

func httpError (err error, w http.ResponseWriter){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

var apiKey string = ""

func checkKey(w http.ResponseWriter, r *http.Request) bool {
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
		}

		}).Methods("GET")


	fmt.Printf("Listening @ 127.0.0.1:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}



func listFLTorrents(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	category, err := strconv.Atoi(vars["category"])
	if err != nil{
		httpError(err, w)
	}
	page, err := strconv.Atoi(vars["page"])
	if err != nil{
		httpError(err, w)
	}

	if !browse.IsCategory(category){
		httpError(errors.New(fmt.Sprintf("%d is not a cat	egory", category)), w)
	}

	torrents, err := filelist.GetTorrents(category, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	fmt.Println()
	fmt.Println(torrents)

	bytes, err := json.MarshalIndent(torrents, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write(bytes)
}