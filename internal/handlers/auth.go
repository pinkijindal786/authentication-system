package handlers

import (
	"Authentication_System/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService services.IAuthService
}

func NewAuthHandler(authService services.IAuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// SignUp handles user registration
func (h *AuthHandler) SignUp(c *gin.Context) {
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
	err := h.AuthService.SignUp(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

// SignIn handles user login and token generation
func (h *AuthHandler) SignIn(c *gin.Context) {
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
	token, err := h.AuthService.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// RefreshToken handles token renewal
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer to refresh the token
	newToken, err := h.AuthService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newToken})
}

// RevokeToken handles token revocation
func (h *AuthHandler) RevokeToken(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Call the service layer to revoke the token
	err := h.AuthService.RevokeToken(req.Token)
	if err != nil {
		if strings.Contains(err.Error(), "token is already revoked") {
			c.JSON(http.StatusOK, gin.H{"message": "token revoked successfully"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token revoked successfully"})
}

// SecureEndpoint is an example of a protected route
func (h *AuthHandler) SecureEndpoint(c *gin.Context) {
	// Retrieve the user ID from the context (set by middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "welcome to the secure endpoint", "user_id": userID})
}
