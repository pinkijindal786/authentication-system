package repositories

import (
	"Authentication_System/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
