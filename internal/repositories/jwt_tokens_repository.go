package repositories

import (
	"Authentication_System/internal/models"
	"strings"

	"gorm.io/gorm"
)

type IJwtTokensRepository interface {
	RevokeToken(token string) error
	IsTokenRevoked(token string) (bool, error)
}

type JwtTokensRepository struct {
	DB *gorm.DB
}

func InitializeJwtTokensRepository(db *gorm.DB) *JwtTokensRepository {
	return &JwtTokensRepository{DB: db}
}

func (r *JwtTokensRepository) RevokeToken(token string) error {
	revokedToken := models.RevokedToken{Token: token}
	return r.DB.Create(&revokedToken).Error
}

func (r *JwtTokensRepository) IsTokenRevoked(token string) (bool, error) {
	var revokedToken models.RevokedToken
	if err := r.DB.Where("token = ?", token).First(&revokedToken).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
