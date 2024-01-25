package repositories

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}

func (repo *userRepo) CreateUser(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) GetUser(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}