package handlers

import (
	"html/template" // however you organize template parsing
	"net/http"
	"romdb-api/internal/dataset"

	"github.com/go-chi/chi/v5"
)

var game_list_tmpl *template.Template
var game_detail_tmpl *template.Template

func SetGameTemplates(list, detail *template.Template) {
	game_list_tmpl = list
	game_detail_tmpl = detail
}

func GameList(w http.ResponseWriter, r *http.Request) {
	d, err := dataset.Load()
	if err != nil {
		http.Error(w, "dataset unavailable", 500)
		return
	}

	type Game struct {
		ID    string
		Title string
	}

	var games []Game
	for id, entry := range d.Entries {
		title, _ := entry["canonical_title"].(string)
		games = append(games, Game{ID: id, Title: title})
	}

	w.Header().Set("Content-Type", "text/html: charset=utf-8")
	if err := game_list_tmpl.Execute(w, games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GameDetail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	d, err := dataset.Load()
	if err != nil {
		http.Error(w, "dataset unavailable", 500)
		return
	}

	entry, ok := d.Entries[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := game_detail_tmpl.Execute(w, entry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
