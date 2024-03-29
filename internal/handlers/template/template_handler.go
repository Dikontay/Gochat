package template

import (
	"net/http"
	"path/filepath"
	"sync"
)
import "html/template"

type TemplateHandler struct {
	once     sync.Once
	Filename string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.Filename)))
	})
	t.templ.Execute(w, r)
}
