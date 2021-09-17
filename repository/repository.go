package repository

import "github.com/wildanadt/go-rest-clean/models"

type UserRepo interface {
	Add(user *models.User) error
	GetUserByUsername(username string) error
	LoginCheck(user *models.User) (models.User, error)
}
