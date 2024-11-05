package controllers

import (
	"github.com/alanzorzi/crud-go/app/model"
	serviceInterface "github.com/alanzorzi/crud-go/app/services/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService serviceInterface.AuthServiceInterface
}

func NewAuthController(service serviceInterface.AuthServiceInterface) *AuthController {
	return &AuthController{authService: service}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	_, token, err := ctrl.authService.LoginUserServices(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
