package models

import "time"

type User struct {
	ID uint
	Name string
	Email string
	Password string
}

const UserMigrationQuery string = "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT)"

// TODO: store browser and ip for improved security
type UserSession struct {
	ID uint
	UserID uint
	Token string
	ExpiresAt time.Time
}

const UserSessionMigrationQuery string = "CREATE TABLE IF NOT EXISTS user_sessions (id INTEGER PRIMARY KEY, user_id INTEGER, token TEXT, expires_at TEXT, FOREIGN KEY(user_id) REFERENCES users(id))"
