package main

import (
	"database/sql"
	"log"

	"github.com/berkaycubuk/mqtt-studio/cmd/api"
	"github.com/berkaycubuk/mqtt-studio/config"
	"github.com/berkaycubuk/mqtt-studio/db"
)

func main() {
	db, err := db.NewSQLiteStorage(config.Envs.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
