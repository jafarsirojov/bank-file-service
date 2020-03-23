package app

import (
	"errors"
	"github.com/jafarsirojov/rest/pkg/rest"
	"net/http"
	"path/filepath"
)

const multipartMaxBytes = 10 * 1024 * 1024

func (s *server) handlePanic() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic(errors.New("some bad things happened"))
	}
}

func (s *server) handleFileUpload() func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		err := request.ParseMultipartForm(multipartMaxBytes)
		if err != nil {
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		files := request.MultipartForm.File["file"]
		type FileURL struct {
			Name string `name`
		}
		fileURLs := make([]FileURL, 0, len(files))
		ext := make(map[string]string)
		ext[".txt"] = contentTypeText
		ext[".pdf"] = contentTypePdf
		ext[".png"] = contentTypePng
		ext[".jpg"] = contentTypeJpg
		ext[".html"] = contentTypeHtml
		for _, file := range files {
			contentType, ok := ext[filepath.Ext(file.Filename)]
			if !ok {
				http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			//contentType := path.Ext(file.Filename)
			openFile, err := file.Open()
			if err != nil {
				http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				continue
			}
			newFile, err := s.filesSvc.Save(openFile, contentType)
			if err != nil {
				http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				continue
			}

			fileURLs = append(fileURLs, FileURL{
				newFile,
			})
		}
		err = rest.WriteJSONBody(responseWriter, fileURLs)
		if err != nil {
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

	}
}
