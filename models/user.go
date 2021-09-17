package models

type User struct {
	Metadata
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
