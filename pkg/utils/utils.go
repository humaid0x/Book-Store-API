package utils

import (
	"time"

	"book-store/pkg/config"
	"book-store/pkg/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ComparePassword(hashedPassword, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func GenerateToken(user *models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix() // Token expires in 15 minutes

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(config.LocalConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}