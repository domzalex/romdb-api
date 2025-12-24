package handlers

import (
	"html/template"
	"net/http"
)

var index_tmpl *template.Template

func SetIndexTemplate(tmpl *template.Template) {
	index_tmpl = tmpl
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := index_tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
