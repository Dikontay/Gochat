package handlers

import (
	"gochat/internal/chat"
	"gochat/internal/handlers/template"
	"net/http"
	"path/filepath"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	// add a css file to route
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("/", &template.TemplateHandler{Filename: "../ui/templates/index.html"})
	r := chat.NewRoom()
	mux.Handle("/room", r)
	go r.Run()
	return mux

}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
