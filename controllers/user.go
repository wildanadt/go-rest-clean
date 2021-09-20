package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/wildanadt/go-rest-clean/usecase"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(r *gin.RouterGroup, uus usecases.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uus,
	}

	r.GET("/test", handler.GetUser)
	// r.POST("/register")
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "api"})
}
