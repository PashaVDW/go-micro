package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// TODO connect to database

	// TODO set up config
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.sprintf(":&s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
