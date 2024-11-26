package repositories

import (
	"authentication_system/internal/models"
	"strings"

	"gorm.io/gorm"
)

type JwtTokensRepository interface {
	RevokeToken(token string) error
	IsTokenRevoked(token string) (bool, error)
}

type JwtTokensRepositoryData struct {
	DB *gorm.DB
}

func InitializeJwtTokensRepository(db *gorm.DB) *JwtTokensRepositoryData {
	return &JwtTokensRepositoryData{DB: db}
}

func (r *JwtTokensRepositoryData) RevokeToken(token string) error {
	revokedToken := models.JWTToken{Token: token}
	return r.DB.Create(&revokedToken).Error
}

func (r *JwtTokensRepositoryData) IsTokenRevoked(token string) (bool, error) {
	var revokedToken models.JWTToken
	if err := r.DB.Where("token = ?", token).First(&revokedToken).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
