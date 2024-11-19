package middlewares

import (
	"Authentication_System/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		c.Abort()
		return
	}

	// Extract the token from the "Bearer <token>" format
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := utils.ValidateJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Optionally extract claims if needed
	claims, err := utils.ExtractClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to extract claims"})
		c.Abort()
		return
	}

	// Set user ID in context for downstream handlers
	c.Set("userID", claims["id"])
	c.Next()
}
