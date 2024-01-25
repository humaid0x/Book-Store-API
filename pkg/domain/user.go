package domain

import (
	"book-store/pkg/models"
)

// for database repository operation (call from service)
type IUserRepo interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
}
