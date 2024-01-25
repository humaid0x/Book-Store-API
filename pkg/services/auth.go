package services

import (
	"errors"

	"book-store/pkg/domain"
	"book-store/pkg/models"
	"book-store/pkg/utils"
)

// authService is the struct implementing IAuthService
type authService struct {
	repo domain.IUserRepo
}

// AuthServiceInstance creates an instance of IAuthService
func AuthServiceInstance(userRepo domain.IUserRepo) domain.IAuthService {
	return &authService{
		repo: userRepo,
	}
}

// Signup creates a new user account
func (service *authService) Signup(user *models.User) error {
	// Check if user already exists
	_, err := service.repo.GetUser(user.UserName)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash the user's password before storing it
	hashedPassword := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	if err := service.repo.CreateUser(user); err != nil {
		return errors.New("user signup failed")
	}
	return nil
}

// Signin verifies user credentials and returns a token on successful signin
func (service *authService) Signin(username, password string) (string, error) {
	user, err := service.repo.GetUser(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", errors.New("error generating token")
	}
	return token, nil
}
