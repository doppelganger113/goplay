package chat

import (
	"sync"
	"net/http"
	"path/filepath"
	"html/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.template.Execute(w, r)
}

func NewTemplateHandler(filename string) *templateHandler {
	return &templateHandler{
		filename: filename,
	}
}
