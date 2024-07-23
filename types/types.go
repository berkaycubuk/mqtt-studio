package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user User) error
}

type User struct {
	ID			int			`json:"id"`
	FullName	string		`json:"fullName"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt	time.Time	`json:"createdAt"`
}

type RegisterUserPayload struct {
	FullName	string	`json:"fullName" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=3,max=130"`
}
