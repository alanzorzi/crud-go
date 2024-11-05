package routes

import (
	"database/sql"
	"github.com/alanzorzi/crud-go/app/middleware"

	"github.com/alanzorzi/crud-go/app/controllers"
	"github.com/alanzorzi/crud-go/app/repository"
	"github.com/alanzorzi/crud-go/app/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	userRepo := repository.NewUserRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	r.POST("/login", authController.Login)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/users", userController.GetAllUsers)
		authRoutes.GET("/user/:id", userController.GetUserById)
		authRoutes.POST("/users", userController.CreateUser)
		authRoutes.PUT("/users/:id", userController.UpdateUser)
		authRoutes.DELETE("/users/:id", userController.DeleteUser)
	}
}
