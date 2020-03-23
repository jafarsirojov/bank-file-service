package main

// package
// import
// var + type
// method + function

import (

	"file-service/cmd/file-service/app"
	"file-service/pkg/file-service/services/files"
	"flag"
	"github.com/jafarsirojov/mux/pkg/mux"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9997", "Server port")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	start(addr)
	log.Println(addr)
}

func start(addr string) {
	router := mux.NewExactMux()

	templatesPath := filepath.Join("web", "templates")
	assetsPath := filepath.Join("web", "assets")
	mediaPath := filepath.Join("web", "media")

	filesSvc := files.NewFilesSvc(mediaPath)
	server := app.NewServer(
		router,
		filesSvc,
		templatesPath,
		assetsPath,
		mediaPath,
	)
	server.InitRoutes()

	panic(http.ListenAndServe(addr, server))
}
