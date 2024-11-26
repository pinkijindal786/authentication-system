package repositories

import (
	"authentication_system/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryData struct {
	DB *gorm.DB
}

func InitializeUserRepository(db *gorm.DB) *UserRepositoryData {
	return &UserRepositoryData{DB: db}
}

func (repo *UserRepositoryData) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepositoryData) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
