package services

import (
	"authentication_system/internal/models"
	"authentication_system/internal/repositories"
	"authentication_system/internal/utils"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUp(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	mockTokenRepo := new(repositories.MockJwtTokensRepository)
	mockUtils := new(utils.MockUtils)

	mockRepo.On("CreateUser", mock.Anything).Return(nil)
	mockUtils.On("HashPassword", mock.Anything).Return("mock", nil)

	authService := InitializeAuthService(mockRepo, mockTokenRepo, mockUtils)

	err := authService.SignUp("test@example.com", "password123")

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSignIn_Success(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	mockTokenRepo := new(repositories.MockJwtTokensRepository)
	mockUtils := new(utils.MockUtils)

	mockRepo.On("GetUserByEmail", "test@example.com").Return(&models.User{
		ID:       1,
		Email:    "test@example.com",
		Password: "hashedPassword",
	}, nil)
	mockUtils.On("GenerateAuthToken", mock.Anything).Return("mockAuthToken", nil)
	mockUtils.On("GenerateRefreshToken", mock.Anything).Return("mockRefreshToken", nil)
	mockUtils.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(true)

	authService := InitializeAuthService(mockRepo, mockTokenRepo, mockUtils)

	token, err := authService.SignIn("test@example.com", "password123")

	assert.Nil(t, err)
	assert.Equal(t, "mockAuthToken", token.AuthToken)
	assert.Equal(t, "mockRefreshToken", token.RefreshToken)
	mockRepo.AssertExpectations(t)
}

func TestRevokeToken_Success(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	mockTokenRepo := new(repositories.MockJwtTokensRepository)
	mockUtils := new(utils.MockUtils)

	authService := InitializeAuthService(mockRepo, mockTokenRepo, mockUtils)

	mockTokenRepo.On("IsTokenRevoked", "validToken").Return(false, nil)
	mockTokenRepo.On("RevokeToken", "validToken").Return(nil)

	err := authService.RevokeToken("validToken")

	assert.Nil(t, err)
	mockTokenRepo.AssertExpectations(t)
}

func TestRevokeToken_AlreadyRevoked(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	mockTokenRepo := new(repositories.MockJwtTokensRepository)
	mockUtils := new(utils.MockUtils)

	authService := InitializeAuthService(mockRepo, mockTokenRepo, mockUtils)

	mockTokenRepo.On("IsTokenRevoked", "revokedToken").Return(true, nil)

	err := authService.RevokeToken("revokedToken")

	assert.NotNil(t, err)
	assert.Equal(t, "token is already revoked", err.Error())
	mockTokenRepo.AssertExpectations(t)
}

func TestRefreshToken_Success(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	mockTokenRepo := new(repositories.MockJwtTokensRepository)
	mockUtils := new(utils.MockUtils)

	authService := InitializeAuthService(mockRepo, mockTokenRepo, mockUtils)

	mockTokenRepo.On("IsTokenRevoked", "oldToken").Return(false, nil)
	mockUtils.On("ValidateJWT", "oldToken").Return(&jwt.Token{
		Raw:    "mockedRawToken",
		Method: jwt.SigningMethodHS256,
		Header: map[string]interface{}{
			"alg": "HS256",
		},
		Claims: jwt.MapClaims{
			"userId": float64(123),
			"exp":    jwt.TimeFunc().Add(time.Hour * 24).Unix(), // 24 hours expiration
		},
		Signature: "mockedSignature",
		Valid:     true, // Ensures the token is marked valid
	}, nil)
	mockUtils.On("GenerateAuthToken", mock.Anything).Return("mockAuthToken", nil)

	token, err := authService.RefreshToken("oldToken")

	assert.Nil(t, err)
	assert.Equal(t, "mockAuthToken", token)
	mockTokenRepo.AssertExpectations(t)
}

func TestRefreshToken_InvalidToken(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	mockTokenRepo := new(repositories.MockJwtTokensRepository)
	mockUtils := new(utils.MockUtils)

	mockTokenRepo.On("IsTokenRevoked", "invalidToken").Return(false, nil)
	mockUtils.On("ValidateJWT", "invalidToken").Return(nil, errors.New("Invalid token"))

	authService := InitializeAuthService(mockRepo, mockTokenRepo, mockUtils)

	_, err := authService.RefreshToken("invalidToken")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid token", err.Error())
}
