package models

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binging:"required"`
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}

type UserInput struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}
