package domain

import (
	"book-store/pkg/models"
)

// for service operation (response to controller | call from controller)
type IAuthService interface {
	Signup(user *models.User) error
	Signin(username, password string) (string, error)
}