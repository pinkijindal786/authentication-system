package handlers

import (
	"Authentication_System/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dependency injection of AuthService
var authService *services.AuthService

// InitializeAuthService initializes the AuthService instance
func InitializeAuthService(service *services.AuthService) {
	authService = service
}

// SignUp handles user registration
func SignUp(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer to register the user
	err := authService.SignUp(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

// SignIn handles user login and token generation
func SignIn(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer to authenticate the user
	token, err := authService.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// RefreshToken handles token renewal
func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer to refresh the token
	newToken, err := authService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newToken})
}

// RevokeToken handles token revocation
func RevokeToken(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Call the service layer to revoke the token
	err := authService.RevokeToken(req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not revoke token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token revoked successfully"})
}

// SecureEndpoint is an example of a protected route
func SecureEndpoint(c *gin.Context) {
	// Retrieve the user ID from the context (set by middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "welcome to the secure endpoint", "user_id": userID})
}
