package routes

import (
	"database/sql"

	"github.com/alanzorzi/crud-go/app/controllers"
	"github.com/alanzorzi/crud-go/app/repository"
	"github.com/alanzorzi/crud-go/app/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	userRepo := repository.NewUserRepository(db)

	// Injetar a service
	userService := services.NewUserService(userRepo)

	// Injetar a service no controller
	userController := controllers.NewUserController(userService)

	r.GET("/users", userController.GetAllUsers)
	r.GET("/user/:id", userController.GetUserById)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:id", userController.UpdateUser)
	r.DELETE("/deleteUser/:id", userController.DeleteUser)
}
