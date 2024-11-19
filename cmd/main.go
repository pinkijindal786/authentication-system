package main

import (
	"Authentication_System/internal/middlewares"
	"Authentication_System/internal/utils"

	"Authentication_System/internal/database"
	"Authentication_System/internal/handlers"
	"Authentication_System/internal/repositories"
	"Authentication_System/internal/services"

	"github.com/gin-gonic/gin"
)

func InitializeUserHandler() *handlers.AuthHandler {
	databaseConnection := database.ConnectDatabase()
	utils := utils.InitializeUtils()
	userRepo := repositories.InitializeUserRepository(databaseConnection)
	tokenRepo := repositories.InitializeJwtTokensRepository(databaseConnection)
	userService := services.InitializeAuthService(userRepo, tokenRepo, utils)
	userHandler := handlers.NewAuthHandler(userService)
	return userHandler
}

func main() {

	r := gin.Default()
	handlers := InitializeUserHandler()

	// Authentication routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.SignUp)
		auth.POST("/signin", handlers.SignIn)
		auth.POST("/refresh", handlers.RefreshToken)
		auth.POST("/revoke", middlewares.AuthMiddleware, handlers.RevokeToken)
	}

	// Protected route example
	r.GET("/secure", middlewares.AuthMiddleware, handlers.SecureEndpoint)

	r.Run(":8080") // Start the server
}
