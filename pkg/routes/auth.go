package routes

import (
	"book-store/pkg/controllers"

	"github.com/labstack/echo/v4"
)

type authRoutes struct {
	echo      *echo.Echo
	authCtr controllers.AuthController
}

func AuthRoutes(echo *echo.Echo, authCtr controllers.AuthController) *authRoutes {
	return &authRoutes{
		echo:      echo,
		authCtr: authCtr,
	}
}

func (ac *authRoutes) InitAuthRoute() {
	e := ac.echo
	ac.initAuthRoutes(e)
}

func (ac *authRoutes) initAuthRoutes(e *echo.Echo) {
	auth := e.Group("/auth")

	auth.POST("/signup", ac.authCtr.Signup)
	auth.POST("/signin", ac.authCtr.Signin)
}
