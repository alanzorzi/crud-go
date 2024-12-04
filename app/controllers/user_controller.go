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

func NewUserController(service serviceInterface.UserServiceInterface) *userController {
	return &userController{userService: service}
}

// GetUserById retrieves user information based on the provided user ID.
// @Summary Find User by ID
// @Description Retrieves user details based on the user ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} string "User information retrieved successfully"
// @Failure 400 {object} string "Error: Invalid user ID"
// @Failure 404 {object} string "User not found"
// @Router /user/{userId} [get]
func (ctrl *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	users, err := ctrl.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetAllUsers retrieves user information based on the provided user ID.
// @Summary Find Users
// @Description Retrieves user details.
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} string "User information retrieved successfully"
// @Failure 400
// @Router /users [get]
func (ctrl *userController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param userRequest body model.User true "User information for registration"
// @Success 200 {object} model.User
// @Failure 400
// @Failure 500
// @Router /users [post]
func (ctrl *userController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.ValidateUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body model.User true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/{userId} [put]
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
	userID.Password = user.Password
	userID.Age = user.Age

	if err := ctrl.userService.UpdateUser(&userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400
// @Failure 500
// @Router /users/{userId} [delete]
func (ctrl *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
