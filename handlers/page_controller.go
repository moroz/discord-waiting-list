package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/moroz/discord-waiting-list/db/queries"
	"github.com/moroz/discord-waiting-list/templates"
)

type pageController struct {
	db *sql.DB
}

func PageController(db *sql.DB) *pageController {
	return &pageController{db}
}

func (p *pageController) Index(w http.ResponseWriter, r *http.Request) {
	// fetch count of existing waiters
	// render the page

	count, err := queries.New(p.db).CountWaiters(r.Context())
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	templates.IndexTemplate.Execute(w, count)
}
