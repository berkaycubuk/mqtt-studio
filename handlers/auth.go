package handlers

import (
	"database/sql"
	//"fmt"
	"net/http"

	"github.com/berkaycubuk/mqtt-studio/views/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h AuthHandler) LoginGET(c echo.Context) error {
	old := make(map[string]string)
	errors := make(map[string]string)

	return render(c, auth.Login(old, errors))
}

func (h AuthHandler) LoginPOST(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	errors := make(map[string]string)
	old := make(map[string]string)
	old["email"] = email
	old["password"] = password

	if email == "" {
		errors["email"] = "Email is required"
	}

	if password == "" {
		errors["password"] = "Password is required"
	}

	// TODO: check is user exists
	/*
	_, err := h.DB.Exec("INSERT INTO projects (name) VALUES (?)", name)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}
	*/

	// TODO: if user exists create new session (maybe and expire older one)
	// and send the token with cookie

	if len(errors) != 0 {
		return render(c, auth.Login(old, errors))
	}

	return c.String(http.StatusCreated, email + " " + password)
}

func (h AuthHandler) RegisterGET(c echo.Context) error {
	old := make(map[string]string)
	errors := make(map[string]string)

	return render(c, auth.Register(old, errors))
}

func (h AuthHandler) RegisterPOST(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	errors := make(map[string]string)
	old := make(map[string]string)
	old["name"] = name
	old["email"] = email
	old["password"] = password

	if name == "" {
		errors["name"] = "Name is required"
	}

	if email == "" {
		errors["email"] = "Email is required"
	}

	if password == "" {
		errors["password"] = "Password is required"
	}

	/*
	_, err := h.DB.Exec("INSERT INTO projects (name) VALUES (?)", name)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}
	*/

	if len(errors) != 0 {
		return render(c, auth.Register(old, errors))
	}

	return c.String(http.StatusCreated, email + " " + password)
}
