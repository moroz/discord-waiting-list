package handlers

import (
	"database/sql"
	"net/http"

	"github.com/moroz/discord-waiting-list/templates"
)

type waiterController struct {
	db *sql.DB
}

func WaiterController(db *sql.DB) *waiterController {
	return &waiterController{db}
}

func (c *waiterController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	templates.FormTemplate.Execute(w, nil)
}
