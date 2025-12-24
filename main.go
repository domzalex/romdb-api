package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"romdb-api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	templates_dir := "templates"
	game_list_tmpl := template.Must(template.ParseFiles(filepath.Join(templates_dir, "game_list.html")))
	game_detail_tmpl := template.Must(template.ParseFiles(filepath.Join(templates_dir, "game_detail.html")))
	index_tmpl := template.Must(template.ParseFiles(filepath.Join(templates_dir, "index.html")))

	handlers.SetIndexTemplate(index_tmpl)
	handlers.SetGameTemplates(game_list_tmpl, game_detail_tmpl)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Get("/", handlers.Index)
	r.Get("/games", handlers.GameList)
	r.Get("/games/{id}", handlers.GameDetail)

	log.Println("Server running on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
