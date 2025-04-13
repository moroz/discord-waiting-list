package handlers

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

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
}

func (c *waiterController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	err := templates.FormTemplate.Execute(w, FormParams{countries})
	if err != nil {
		log.Print(err)
	}
}

type SignUpParams struct {
	Email  string
	Name   string
	Bio    *string
	Region *string
}

type RegionOption struct {
	Label string
	Value string
}

func (p *SignUpParams) Parse(f url.Values) {
	p.Email = f.Get("email")
	p.Name = f.Get("name")
	if b := f.Get("bio"); b != "" {
		p.Bio = &b
	}
	if r := f.Get("region"); r != "" {
		p.Region = &r
	}
}

func (c *waiterController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var params SignUpParams
	params.Parse(r.PostForm)

	log.Printf("%#v", params)
}
