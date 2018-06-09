package main

import (
	"github.com/teo-mateo/flbrowser/filelist/server"
	"flag"
	"github.com/teo-mateo/flbrowser/filelist/rtorrent"
	"path/filepath"
	"os"
	"fmt"
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

	flag.IntVar(&port, "port", 8080, "-port=8080")
	flag.StringVar(&apikey, "apikey", "abcdefgh", "--apikey=abcdefgh")
	flag.StringVar(&rtorrent.Ru, "ru", "nouser", "--ru=username")
	flag.StringVar(&rtorrent.Rp, "rp", "nopwd", "--rp=password")
	flag.StringVar(&rtorrent.RActive, "ractive", cwd, "--ractive=/path/to/.torrent/files")
	flag.StringVar(&rtorrent.RDownloads, "rdown", cwd, "--rdown=/path/to/downloads")
	flag.StringVar(&rtorrent.RSession, "rsess", cwd, "--rsess=/path/to/session")
	flag.Parse()

	fmt.Println("Directories:::::::")
	fmt.Printf("rActive:%s\n", rtorrent.RActive)
	fmt.Printf("rDownloads:%s\n", rtorrent.RDownloads)
	fmt.Printf("rSession:%s\n", rtorrent.RSession)

	server.Start(port, apikey)
}
