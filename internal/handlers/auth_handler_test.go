package handlers_test

import (
	"authentication_system/internal/handlers"
	"authentication_system/internal/models"
	"authentication_system/internal/services"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

// TestSignUp_Success tests successful user registration
func TestSignUp_Success(t *testing.T) {
	mockService := &services.MockAuthService{}
	mockService.On("SignUp", "test@example.com", "password123").Return(nil)

	handler := handlers.NewAuthHandler(mockService)
	router := setupRouter()
	router.POST("/signup", handler.SignUp)

	body := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.JSONEq(t, `{"message":"user created successfully"}`, resp.Body.String())
	mockService.AssertCalled(t, "SignUp", "test@example.com", "password123")
}

// TestSignUp_Failure tests user registration failure
func TestSignUp_Failure(t *testing.T) {
	mockService := &services.MockAuthService{}
	mockService.On("SignUp", "test@example.com", "password123").Return(errors.New("some error"))

	handler := handlers.NewAuthHandler(mockService)
	router := setupRouter()
	router.POST("/signup", handler.SignUp)

	body := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.JSONEq(t, `{"error":"could not create user"}`, resp.Body.String())
	mockService.AssertCalled(t, "SignUp", "test@example.com", "password123")
}

// TestSignIn_Success tests successful user login
func TestSignIn_Success(t *testing.T) {
	mockService := &services.MockAuthService{}
	mockService.On("SignIn", "test@example.com", "password123").Return(&models.SignInResponse{
		AuthToken:    "mockToken",
		RefreshToken: "mockToken",
	}, nil)

	handler := handlers.NewAuthHandler(mockService)
	router := setupRouter()
	router.POST("/signin", handler.SignIn)

	body := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"authToken":"mockToken","refreshToken":"mockToken"}`, resp.Body.String())
	mockService.AssertCalled(t, "SignIn", "test@example.com", "password123")
}

// TestSignIn_InvalidCredentials tests login with invalid credentials
func TestSignIn_InvalidCredentials(t *testing.T) {
	mockService := &services.MockAuthService{}
	mockService.On("SignIn", "test@example.com", "wrongpassword").Return(nil, errors.New("invalid credentials"))

	handler := handlers.NewAuthHandler(mockService)
	router := setupRouter()
	router.POST("/signin", handler.SignIn)

	body := `{"email":"test@example.com","password":"wrongpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.JSONEq(t, `{"error":"invalid credentials"}`, resp.Body.String())
	mockService.AssertCalled(t, "SignIn", "test@example.com", "wrongpassword")
}

// TestRefreshToken_Success tests successful token refresh
func TestRefreshToken_Success(t *testing.T) {
	mockService := &services.MockAuthService{}
	mockService.On("RefreshToken", "mockRefreshToken").Return("newMockToken", nil)

	handler := handlers.NewAuthHandler(mockService)
	router := setupRouter()
	router.POST("/refresh", handler.RefreshToken)

	body := `{"refresh_token":"mockRefreshToken"}`
	req, _ := http.NewRequest(http.MethodPost, "/refresh", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"access_token":"newMockToken"}`, resp.Body.String())
	mockService.AssertCalled(t, "RefreshToken", "mockRefreshToken")
}

// TestRevokeToken_Success tests successful token revocation
func TestRevokeToken_Success(t *testing.T) {
	mockService := &services.MockAuthService{}
	mockService.On("RevokeToken", "mockToken").Return(nil)

	handler := handlers.NewAuthHandler(mockService)
	router := setupRouter()
	router.POST("/revoke", handler.RevokeToken)

	body := `{"token":"mockToken"}`
	req, _ := http.NewRequest(http.MethodPost, "/revoke", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"message":"token revoked successfully"}`, resp.Body.String())
	mockService.AssertCalled(t, "RevokeToken", "mockToken")
}
