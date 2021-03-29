package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static
var embededFiles embed.FS

var useLive bool

func main() {
	flag.BoolVar(&useLive, "live", true, "is use live mode(default:true)")
	flag.Parse()
	log.Print("live:", useLive)
	http.Handle("/", http.FileServer(getFileSystem(useLive)))
	http.ListenAndServe(":8888", nil)
}

func getFileSystem(useLive bool) http.FileSystem {
	if useLive {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
