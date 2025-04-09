package handlers

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	pages := PageController(db)
	r.Get("/", pages.Index)

	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	return r
}
