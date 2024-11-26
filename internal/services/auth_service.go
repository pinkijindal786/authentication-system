package services

import (
	"authentication_system/internal/models"
	"authentication_system/internal/repositories"
	"authentication_system/internal/utils"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	SignUp(email, password string) error
	SignIn(email, password string) (*models.SignInResponse, error)
	RevokeToken(token string) error
	RefreshToken(refreshToken string) (string, error)
}

type AuthServiceData struct {
	Repo      repositories.UserRepository
	TokenRepo repositories.JwtTokensRepository
	Utils     utils.Utils
}

func InitializeAuthService(repo repositories.UserRepository, tokenRepo repositories.JwtTokensRepository, utils utils.Utils) *AuthServiceData {
	return &AuthServiceData{Repo: repo, TokenRepo: tokenRepo, Utils: utils}
}

func (s *AuthServiceData) SignUp(email, password string) error {
	hashedPassword, err := s.Utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: hashedPassword,
		IsActive: true,
	}

	return s.Repo.CreateUser(user)
}

func (s *AuthServiceData) SignIn(email, password string) (*models.SignInResponse, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil || !s.Utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	authToken, err := s.Utils.GenerateAuthToken(user.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.Utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &models.SignInResponse{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceData) RevokeToken(token string) error {
	isRevoked, err := s.TokenRepo.IsTokenRevoked(token)
	if err != nil {
		fmt.Printf("Failed to revoke the token %s", err)
		return err
	}
	if isRevoked {
		return errors.New("token is already revoked")
	}
	return s.TokenRepo.RevokeToken(token)
}

func (s *AuthServiceData) RefreshToken(refreshToken string) (string, error) {
	// Check if token is revoked
	isRevoked, err := s.TokenRepo.IsTokenRevoked(refreshToken)
	if err != nil {
		return "", err
	}
	if isRevoked {
		return "", errors.New("token is revoked")
	}

	// Validate and extract userID from old token
	parsedToken, err := s.Utils.ValidateJWT(refreshToken)
	if err != nil {
		return "", errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid token claims")
	}
	userID := uint(claims["userId"].(float64))

	// Generate new auth token
	newAuthToken, err := s.Utils.GenerateAuthToken(userID)
	if err != nil {
		return "", err
	}

	return newAuthToken, nil
}
