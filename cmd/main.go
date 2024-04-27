package main

import (
	"database/sql"
	"fmt"

	"github.com/berkaycubuk/mqtt-studio/handlers"
	"github.com/berkaycubuk/mqtt-studio/models"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func runMigrations(db *sql.DB) {
	// projects
	projectsMigration, err := db.Prepare(models.ProjectMigrationQuery)
	if err != nil {
		fmt.Println("Error occured on Projects migration: ", err.Error())
	}
	projectsMigration.Exec()
}

func main() {
	fmt.Println("MQTT Studio v2")

	// database
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	runMigrations(db)

	defer db.Close()

	e := echo.New()

	projectHandler := handlers.ProjectHandler{
		DB: db,
	}

	e.GET("/projects", projectHandler.Index)
	e.GET("/projects/create", projectHandler.Create)
	e.POST("/projects/create", projectHandler.CreatePost)
	e.Logger.Fatal(e.Start(":2000"))
}
