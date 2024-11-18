package services

import (
	"Authentication_System/models"
	"Authentication_System/repositories"
	"Authentication_System/utils"
	"errors"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	Repo      *repositories.UserRepository
	TokenRepo *repositories.JwtTokensRepository
}

func (s *AuthService) SignUp(email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
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

func (s *AuthService) SignIn(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(user.ID)
}

func (s *AuthService) RevokeToken(token string) error {
	isRevoked, err := s.TokenRepo.IsTokenRevoked(token)
	if err != nil {
		return err
	}
	if isRevoked {
		return errors.New("token is already revoked")
	}
	return s.TokenRepo.RevokeToken(token)
}

func (s *AuthService) RefreshToken(oldToken string) (string, error) {
	// Check if token is revoked
	isRevoked, err := s.TokenRepo.IsTokenRevoked(oldToken)
	if err != nil {
		return "", err
	}
	if isRevoked {
		return "", errors.New("token is revoked")
	}

	// Validate and extract userID from old token
	parsedToken, err := utils.ValidateJWT(oldToken)
	if err != nil {
		return "", errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid token claims")
	}
	// email := claims["email"].(string)
	userID := uint(claims["userId"].(uint))

	// Generate new token
	newToken, err := utils.GenerateJWT(userID)
	if err != nil {
		return "", err
	}

	// Revoke old token
	if err := s.TokenRepo.RevokeToken(oldToken); err != nil {
		return "", err
	}

	return newToken, nil
}
