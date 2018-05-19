package main

import (
	"github.com/teo-mateo/flbrowser/filelist/server"
	"flag"
	"github.com/teo-mateo/flbrowser/filelist/rtorrent"
	"log"
	"github.com/davecgh/go-spew/spew"
)

var port int
var apikey string
var ru string
var rp string

func main() {


	flag.IntVar(&port, "port", 8080, "-port=8080")
	flag.StringVar(&apikey, "apikey", "abcdefgh", "--apikey=abcdefgh")
	flag.StringVar(&rtorrent.Ru, "ru", "nouser", "--ru=username")
	flag.StringVar(&rtorrent.Rp, "rp", "nopwd", "--rp=password")
	flag.StringVar(&rtorrent.RActive, "ractive", "", "--ractive=/path/to/.torrent/files")
	flag.StringVar(&rtorrent.RDownloads, "rdown", "", "--rdown=/path/to/downloads")
	flag.StringVar(&rtorrent.RSession, "rsess", "", "--rsess=/path/to/session")

	flag.Parse()

	server.Start(port, apikey)



}


func tests(){
	list, err := rtorrent.GetTorrents()
	if err != nil{
		log.Panic(err)
	}

	spew.Dump(list)

}
