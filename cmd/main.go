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
	fmt.Println("System: Running migrations.")

	// projects
	projectsMigration, err := db.Prepare(models.ProjectMigrationQuery)
	if err != nil {
		fmt.Println("Error occured on Projects migration: ", err.Error())
	}
	projectsMigration.Exec()
	fmt.Println("System: Projects migration ran.")

	// users
	usersMigration, err := db.Prepare(models.UserMigrationQuery)
	if err != nil {
		fmt.Println("Error occured on Users migration: ", err.Error())
	}
	usersMigration.Exec()
	fmt.Println("System: Users migration ran.")

	// user_sessions
	userSessionsMigration, err := db.Prepare(models.UserSessionMigrationQuery)
	if err != nil {
		fmt.Println("Error occured on User Sessions migration: ", err.Error())
	}
	userSessionsMigration.Exec()
	fmt.Println("System: User Sessions migration ran.")
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
	landingHandler := handlers.LandingHandler{}
	authHandler := handlers.AuthHandler{
		DB: db,
	}

	e.Static("/", "public")
	e.GET("/", landingHandler.Index)
	e.GET("/login", authHandler.LoginGET)
	e.POST("/login", authHandler.LoginPOST)
	e.GET("/register", authHandler.RegisterGET)
	e.POST("/register", authHandler.RegisterPOST)
	e.GET("/projects", projectHandler.Index)
	e.GET("/projects/create", projectHandler.Create)
	e.POST("/projects/create", projectHandler.CreatePost)
	e.Logger.Fatal(e.Start(":2000"))
}
