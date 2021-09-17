package usecases

import (
	"fmt"

	"github.com/wildanadt/go-rest-clean/models"
	"github.com/wildanadt/go-rest-clean/repository"
)

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserHandler(r repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: r,
	}
}

func (uus *userUsecase) AddUser(input *models.RegisterInput) error {
	var user models.User
	user.Username = input.Username
	user.Password = input.Password
	err := uus.userRepo.Add(&user)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (uus *userUsecase) CheckUser(input *models.LoginInput) error {
	err := uus.userRepo.GetUserByUsername(input.Username)
	if err != nil {
		return err
	}

	return nil
}
