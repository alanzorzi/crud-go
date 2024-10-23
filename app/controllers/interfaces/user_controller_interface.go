package interfaces

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	GetUserById(c *gin.Context)
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
