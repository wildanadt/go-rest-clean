package repository

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/wildanadt/go-rest-clean/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	db.AutoMigrate(&models.User{})

	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Add(user *models.User) error {

	//check user existing
	err := u.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}

	//hashing user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	err = u.db.Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (u *userRepo) GetUserByUsername(username string) error {
	user := models.User{}
	user.Username = username
	result := u.db.Find(&user)
	// fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		err := errors.New("user exist")
		return err
	}
	return nil
}

func (u *userRepo) LoginCheck(user *models.User) (models.User, error) {

	_user := models.User{}

	err := u.db.Model(models.User{}).Where("username = ?", user.Username).Take(&_user).Error
	if err != nil {
		return _user, err
	}

	return _user, err

}
