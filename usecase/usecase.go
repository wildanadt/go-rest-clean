package usecases

import "github.com/wildanadt/go-rest-clean/models"

type UserUsecase interface {
	AddUser(user *models.RegisterInput) error
	CheckUser(input *models.LoginInput) error
}

type AuthUsecase interface {
	LoginCheck(username string, password string) (string, error)
}
