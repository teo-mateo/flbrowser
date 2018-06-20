package main

import (
	"github.com/teo-mateo/flbrowser/filelist/server"
	"flag"
	"github.com/teo-mateo/flbrowser/filelist/rtorrent"
	"path/filepath"
	"os"
	"fmt"
	"log"
)

var port int
var apikey string
var httpUser string
var httpPwd string
var ru string
var rp string
var clientDir string


func main() {

	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil{
		panic(err)
	}

	flag.IntVar(&port, "port", 8080, "-port=8080")
	flag.StringVar(&apikey, "apikey", "abcdefgh", "--apikey=abcdefgh")
	flag.StringVar(&httpUser, "httpu", "", "--httpu=username")
	flag.StringVar(&httpPwd, "httpp", "", "--httpp=password")
	flag.StringVar(&clientDir, "clientdir", filepath.Join(cwd, "../client/dist"), "--clientdir=/path/to/client")
	flag.StringVar(&rtorrent.Ru, "ru", "nouser", "--ru=username")
	flag.StringVar(&rtorrent.Rp, "rp", "nopwd", "--rp=password")
	flag.StringVar(&rtorrent.RActive, "ractive", cwd, "--ractive=/path/to/.torrent/files")
	flag.StringVar(&rtorrent.RDownloads, "rdown", cwd, "--rdown=/path/to/downloads")
	flag.StringVar(&rtorrent.RSession, "rsess", cwd, "--rsess=/path/to/session")
	flag.StringVar(&rtorrent.Raddress, "raddr", "h.bardici.ro:8008/RPC2", "--raddrr=server:port/RPC2")
	flag.Parse()

	fmt.Println()
	fmt.Printf("HTTP Basic Auth Username: %s\n", httpUser)
	fmt.Printf("HTTP Basic Auth Password: %s\n", httpPwd)
	fmt.Println("\nDirectories: ")
	fmt.Printf("rActive:%s\n", rtorrent.RActive)
	fmt.Printf("rDownloads:%s\n", rtorrent.RDownloads)
	fmt.Printf("rSession:%s\n", rtorrent.RSession)
	fmt.Printf("RTR server:%s\n", rtorrent.Raddress)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("current dir: %s\n", dir)
	fmt.Printf("client dir: %s\n", clientDir)

	server.Start(port, apikey, httpUser, httpPwd, clientDir)


}
