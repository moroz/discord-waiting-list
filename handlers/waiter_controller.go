package handlers

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/moroz/discord-waiting-list/services"
	"github.com/moroz/discord-waiting-list/templates"
)

//go:embed countries.json
var countriesJSON []byte

var countries []RegionOption

func init() {
	json.Unmarshal(countriesJSON, &countries)
}

type waiterController struct {
	db *sql.DB
}

func WaiterController(db *sql.DB) *waiterController {
	return &waiterController{db}
}

type FormParams struct {
	Countries []RegionOption
	Params    services.SignUpParams
	Errors    services.ValidationErrors
}

func (c *waiterController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	err := templates.FormTemplate.Execute(w, FormParams{
		Countries: countries,
		Params:    services.SignUpParams{},
		Errors:    nil,
	})
	if err != nil {
		log.Print(err)
	}
}

type RegionOption struct {
	Label string
	Value string
}

func (c *waiterController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var params services.SignUpParams
	params.Parse(r.PostForm)

	_, err = services.WaiterService(c.db).CreateWaiter(r.Context(), &params)

	if err, ok := err.(services.ValidationErrors); ok {
		templates.FormTemplate.Execute(w, FormParams{
			Countries: countries,
			Params:    params,
			Errors:    err,
		})
	}

	if err == nil {
		http.Redirect(w, r, "/success", http.StatusFound)
		return
	}

	log.Print(err)
}

func (c *waiterController) Success(w http.ResponseWriter, r *http.Request) {
	err := templates.SuccessTemplate.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
