package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wildanadt/go-rest-clean/models"
	usecases "github.com/wildanadt/go-rest-clean/usecase"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(r *gin.RouterGroup, uus usecases.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uus,
	}

	r.POST("/register", handler.AddUser)
	// r.POST("/register")
}

func (uh *UserHandler) AddUser(c *gin.Context) {
	var user models.RegisterInput

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = uh.UserUsecase.AddUser(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
