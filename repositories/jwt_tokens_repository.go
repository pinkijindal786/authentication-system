package repositories

import (
	"Authentication_System/models"

	"gorm.io/gorm"
)

type JwtTokensRepository struct {
	DB *gorm.DB
}

func (r *JwtTokensRepository) RevokeToken(token string) error {
	revokedToken := models.RevokedToken{Token: token}
	return r.DB.Create(&revokedToken).Error
}

func (r *JwtTokensRepository) IsTokenRevoked(token string) (bool, error) {
	var revokedToken models.RevokedToken
	if err := r.DB.Where("token = ?", token).First(&revokedToken).Error; err != nil {
		return false, err
	}
	return true, nil
}
