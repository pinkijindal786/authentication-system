package utils

import (
	"authentication_system/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Utils interface {
	GenerateAuthToken(userID uint) (string, error)
	GenerateRefreshToken(userID uint) (string, error)
	HashPassword(password string) (string, error)
	ValidateJWT(token string) (*jwt.Token, error)
	CheckPasswordHash(password, hash string) bool
	ExtractClaims(token *jwt.Token) (jwt.MapClaims, error)
}

type UtilsData struct{}

func InitializeUtils() *UtilsData {
	return &UtilsData{}
}

// HashPassword generates a bcrypt hash of the given password.
func (s *UtilsData) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
func (s *UtilsData) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UtilsData) GenerateAuthToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"type":   "auth",
		"userId": userID,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func (s *UtilsData) GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"type":   "refresh",
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func (s *UtilsData) ExtractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

func (s *UtilsData) ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
}
