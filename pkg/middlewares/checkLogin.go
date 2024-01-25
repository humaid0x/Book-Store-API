package middlewares

import (
	"book-store/pkg/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config := config.LocalConfig
		authHeader := c.Request().Header.Get("Authorization")

		// Check if the Authorization header is provided
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Authorization header missing")
		}

		// Extract the token from the header (Bearer <token>)
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		// Parse the token
		token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}
}
