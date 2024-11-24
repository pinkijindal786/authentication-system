package repositories

import (
	"Authentication_System/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateUser_Success tests that a user is created successfully
func TestCreateUser_Success(t *testing.T) {
	repo := InitializeUserRepository(cleanup("TestCreateUser_Success"))

	user := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
	}
	err := repo.CreateUser(user)

	assert.Nil(t, err)
}

// TestCreateUser_Failure tests that an error is returned if the user creation fails
func TestCreateUser_Failure(t *testing.T) {
	repo := InitializeUserRepository(cleanup("TestCreateUser_Failure"))

	user := &models.User{
		Email:    "test@gmail.com",
		Password: "hashedpassword",
	}
	err := repo.CreateUser(user)

	assert.NotNil(t, err)
}

// TestGetUserByEmail_Success tests that a user is retrieved successfully by email
func TestGetUserByEmail_Success(t *testing.T) {
	repo := InitializeUserRepository(cleanup("TestGetUserByEmail_Success"))

	user, err := repo.GetUserByEmail("test@gmail.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@gmail.com", user.Email)
}

// TestGetUserByEmail_NotFound tests that an error is returned if the user is not found
func TestGetUserByEmail_NotFound(t *testing.T) {
	repo := InitializeUserRepository(cleanup("TestGetUserByEmail_NotFound"))

	user, err := repo.GetUserByEmail("nonexistent@example.com")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}
