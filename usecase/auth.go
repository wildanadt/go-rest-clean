package usecases

import (
	"fmt"

	"github.com/wildanadt/go-rest-clean/models"
	"github.com/wildanadt/go-rest-clean/repository"
	"github.com/wildanadt/go-rest-clean/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	UserRepo repository.UserRepo
}

func NewAuthHanadler(r repository.UserRepo) AuthUsecase {
	return &authUsecase{
		UserRepo: r,
	}
}

func VerifyPassword(password, hashedPasword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}

func (aus *authUsecase) LoginCheck(username string, password string) (string, error) {
	var user models.User
	user.Username = username
	user.Password = password
	user, err := aus.UserRepo.LoginCheck(&user)
	if err != nil {
		return "", err
	}

	fmt.Println(password)
	fmt.Println(user.Password)
	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	_token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return _token, err

}
