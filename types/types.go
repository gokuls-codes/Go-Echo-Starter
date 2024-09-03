package types

import "time"


type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user *User) error
}

type LoginPayload struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RegisterPayload struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"firstName"`
	Email string`json:"email"`
	Password string`json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}