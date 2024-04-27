package models

type Project struct {
	ID uint
	Name string
}

const ProjectMigrationQuery string = "CREATE TABLE IF NOT EXISTS projects (id INTEGER PRIMARY KEY, name TEXT)"
