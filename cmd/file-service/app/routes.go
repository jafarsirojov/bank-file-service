package app

import (
	"github.com/jafarsirojov/mux/pkg/mux"
	_ "github.com/jafarsirojov/mux/pkg/mux/middleware/authorized"
	"github.com/jafarsirojov/mux/pkg/mux/middleware/logger"
	"github.com/jafarsirojov/mux/pkg/mux/middleware/recoverer"
	"net/http"
)

func (receiver *server) InitRoutes() {
	mux := receiver.router.(*mux.ExactMux)
	mux.GET("/api/panic", receiver.handlePanic(), recoverer.Recoverer(), logger.Logger("DEBUG`"))
	mux.GET("/upload", receiver.handleFileUpload())
	mux.POST("/upload", receiver.handleFileUpload())
	mux.GET("/media/{id}", http.StripPrefix("/media", http.FileServer(http.Dir(receiver.mediaPath))).ServeHTTP)
}
