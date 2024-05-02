package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

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
		errors["name"] = "Name is required."
	}

	if email == "" {
		errors["email"] = "Email is required."
	}

	if password == "" {
		errors["password"] = "Password is required."
	}

	// is email being used by any other account
	row := h.DB.QueryRow("SELECT id, email FROM users WHERE email = ?", email)
	if row.Scan() != sql.ErrNoRows {
		fmt.Println(row.Err())
		errors["email"] = "Email is already in use."
	}

	if len(errors) != 0 {
		return render(c, auth.Register(old, errors))
	}

	// insert user
	insert, err := h.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}

	user_id, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}

	// generate session token
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprint(user_id)))
	token := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	// set expiration time
	// 30 days from now
	expireTime := time.Now().Add(24 * 30 * time.Hour)

	// insert session
	_, err = h.DB.Exec("INSERT INTO user_sessions (user_id, token, expires_at) VALUES (?, ?, ?)", user_id, token, expireTime)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "error")
	}

	// create session cookie
	cookie := new(http.Cookie)
	cookie.Name = "user_token"
	cookie.Value = string(token)
	cookie.Expires = expireTime
	c.SetCookie(cookie)

	// login user without manually logging in
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}
