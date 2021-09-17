package usecases

import (
	"go/token"

	"github.com/wildanadt/go-rest-clean/models"
	"github.com/wildanadt/go-rest-clean/repository"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	UserRepo repository.UserRepo
}

func NewAuthHanadler(r repository.UserRepo) AuthUsecaase {
	return &userUsecase{
		userRepo: r,
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

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}


	token, err := token.Ge

}
