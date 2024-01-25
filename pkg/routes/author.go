package routes

import (
	"book-store/pkg/controllers"
	"book-store/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

type authorRoutes struct {
	echo      *echo.Echo
	authorCtr controllers.AuthorController
}

func AuthorRoutes(echo *echo.Echo, authorCtr controllers.AuthorController) *authorRoutes {
	return &authorRoutes{
		echo:      echo,
		authorCtr: authorCtr,
	}
}

func (ac *authorRoutes) InitAuthorRoute() {
	e := ac.echo
	ac.initAuthorRoutes(e)
}

func (ac *authorRoutes) initAuthorRoutes(e *echo.Echo) {
	author := e.Group("/authors")

	author.POST("/", ac.authorCtr.CreateAuthor, middlewares.CheckLogin)
	author.GET("/", ac.authorCtr.GetAuthor)
	author.PUT("/:authorID", ac.authorCtr.UpdateAuthor, middlewares.CheckLogin)
	author.DELETE("/:authorID", ac.authorCtr.DeleteAuthor, middlewares.CheckLogin)
}
