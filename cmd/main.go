package main

import (
	"authentication_system/internal/middlewares"
	"authentication_system/internal/utils"

	"authentication_system/internal/database"
	"authentication_system/internal/handlers"
	"authentication_system/internal/repositories"
	"authentication_system/internal/services"

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

	r.Run(":8080") // Start the server
}
