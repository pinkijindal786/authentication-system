package repositories

import (
	"authentication_system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RevokeToken_Success tests that a token is revoked successfully
func TestRevokeToken_Success(t *testing.T) {
	repo := InitializeJwtTokensRepository(cleanup("RevokeToken_Success"))

	token := &models.JWTToken{
		Token: "dummy_token",
	}
	err := repo.RevokeToken(token.Token)

	assert.Nil(t, err)
}

func TestIsTokenRevoked_Success(t *testing.T) {
	repo := InitializeJwtTokensRepository(cleanup("TestIsTokenRevoked_Success"))

	token := &models.JWTToken{
		Token: "dummy",
	}
	found, err := repo.IsTokenRevoked(token.Token)

	assert.Nil(t, err)
	assert.Equal(t, found, true)
}

func TestIsTokenRevoked_Failure(t *testing.T) {
	repo := InitializeJwtTokensRepository(cleanup("TestIsTokenRevoked_Failure"))

	token := &models.JWTToken{
		Token: "dummy-test",
	}
	found, err := repo.IsTokenRevoked(token.Token)

	assert.Nil(t, err)
	assert.Equal(t, found, false)
}
