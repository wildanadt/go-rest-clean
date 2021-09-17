package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildanadt/go-rest-clean/models"
	usecases "github.com/wildanadt/go-rest-clean/usecase"
)

type AuthHandler struct {
	UserUsecases usecases.UserUsecase
}

func NewAuthhandler(r *gin.RouterGroup, uus usecases.UserUsecase) {
	handler := &AuthHandler{
		UserUsecases: uus,
	}

	r.POST("/register", handler.Register)
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var input models.RegisterInput

	//validate register input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.UserUsecases.AddUser(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"success": "register succesfull"})
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var input models.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.UserUsecases.CheckUser(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
