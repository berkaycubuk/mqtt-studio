package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/berkaycubuk/mqtt-studio/models"
	"github.com/berkaycubuk/mqtt-studio/views/project"
	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	DB *sql.DB
}

func (h ProjectHandler) Index(c echo.Context) error {
	rows, err := h.DB.Query("SELECT * FROM projects")
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			fmt.Println(err.Error())
		}
		projects = append(projects, p)
	}

	return render(c, project.Index(projects))
}

func (h ProjectHandler) Create(c echo.Context) error {

	return render(c, project.Create())
}

func (h ProjectHandler) CreatePost(c echo.Context) error {
	name := c.FormValue("name")

	_, err := h.DB.Exec("INSERT INTO projects (name) VALUES (?)", name)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}

	return c.String(http.StatusCreated, "hit")
}
