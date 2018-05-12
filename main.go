package main

import (
	"net/http"
	"github.com/teo-mateo/flbrowser/filelist/server"
	"flag"
)

func httpError (err error, w http.ResponseWriter){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func main() {

	port := flag.Int("port", 8080, "--port:8080")
	apikey := flag.String("apikey", "abcdefgh", "--apikey:abcdef")
	flag.Parse()

	server.Start(*port, *apikey)
}
//test
