package main

import (
	"Authentication_System/handlers"
	"Authentication_System/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

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
