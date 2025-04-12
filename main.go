package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/moroz/discord-waiting-list/config"
	"github.com/moroz/discord-waiting-list/handlers"
	_ "modernc.org/sqlite"
)

func init() {
	if !config.LOG_TIMESTAMPS {
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	}
}

func main() {
	db, err := sql.Open("sqlite", config.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to open database file: %s", err)
	}

	r := handlers.Router(db)

	log.Printf("Listening on %s", config.LISTEN_ON)

	log.Fatal(http.ListenAndServe(config.LISTEN_ON, r))
}
