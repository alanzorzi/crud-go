package controllers

import (
	"net/http"

	"github.com/alanzorzi/crud-go/app/model"
	serviceInterface "github.com/alanzorzi/crud-go/app/services/interfaces"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService serviceInterface.UserServiceInterface
}

// NewUserController cria uma nova instância de userController
func NewUserController(service serviceInterface.UserServiceInterface) *userController {
	return &userController{userService: service}
}

func (ctrl *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	users, err := ctrl.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetAllUsers retorna todos os usuários do banco
func (ctrl *userController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser cria um novo usuário
func (ctrl *userController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ctrl *userController) UpdateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userID model.User

	userID.ID = c.Param("id")
	userID.Name = user.Name
	userID.Email = user.Email
	userID.Senha = user.Senha
	userID.Age = user.Age

	if err := ctrl.userService.UpdateUser(&userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ctrl *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
