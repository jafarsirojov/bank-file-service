package app

import (
	"errors"
	"file-service/pkg/file-service/services/files"
	"net/http"
)

// описание сервиса, который хранит зависимости и выполняет работу
type server struct { // <- Alt + Enter -> Constructor
	// зависимости (dependencies)
	router        http.Handler
	filesSvc      *files.FilesSvc
	templatesPath string
	assetsPath    string
	mediaPath     string
}

// Все зависимости делятся на:
// 1. required <-
// 2. optional

// crash early
func NewServer(router http.Handler,  filesSvc *files.FilesSvc, templatesPath string, assetsPath string, mediaPath string) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}

	if filesSvc == nil {
		panic(errors.New("filesSvc can't be nil"))
	}
	if templatesPath == "" {
		panic(errors.New("templatesPath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}
	if mediaPath == "" {
		panic(errors.New("mediaPath can't be empty"))
	}

	return &server{
		router:        router,
		filesSvc:      filesSvc,
		templatesPath: templatesPath,
		assetsPath:    assetsPath,
		mediaPath:     mediaPath,
	}
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}

