package main

import (
	"github.com/teo-mateo/flbrowser/filelist/server"
	"flag"
	"github.com/teo-mateo/flbrowser/filelist/rtorrent"
	"path/filepath"
	"os"
)

var port int
var apikey string
var ru string
var rp string

func main() {

	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil{
		panic(err)
	}

	//test4

	flag.IntVar(&port, "port", 8080, "-port=8080")
	flag.StringVar(&apikey, "apikey", "abcdefgh", "--apikey=abcdefgh")
	flag.StringVar(&rtorrent.Ru, "ru", "nouser", "--ru=username")
	flag.StringVar(&rtorrent.Rp, "rp", "nopwd", "--rp=password")
	flag.StringVar(&rtorrent.RActive, "ractive", cwd, "--ractive=/path/to/.torrent/files")
	flag.StringVar(&rtorrent.RDownloads, "rdown", cwd, "--rdown=/path/to/downloads")
	flag.StringVar(&rtorrent.RSession, "rsess", cwd, "--rsess=/path/to/session")
	flag.Parse()

	server.Start(port, apikey)
}
