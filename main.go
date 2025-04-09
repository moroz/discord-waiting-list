package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/moroz/discord-waiting-list/handlers"
	_ "modernc.org/sqlite"
)

const LISTEN_ON = ":3000"

func main() {
	db, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		log.Fatal(err)
	}

	r := handlers.Router(db)

	log.Printf("Listening on %s", LISTEN_ON)

	log.Fatal(http.ListenAndServe(LISTEN_ON, r))
}
