package routes

import (
	"book-store/pkg/controllers"
	"book-store/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

type bookRoutes struct {
	echo    *echo.Echo
	bookCtr controllers.BookController
}

func BookRoutes(echo *echo.Echo, bookCtr controllers.BookController) *bookRoutes {
	return &bookRoutes{
		echo:    echo,
		bookCtr: bookCtr,
	}
}

func (bc *bookRoutes) InitBookRoute() {
	e := bc.echo
	bc.initBookRoutes(e)
}

func (bc *bookRoutes) initBookRoutes(e *echo.Echo) {
	book := e.Group("/books")

	book.POST("/", bc.bookCtr.CreateBook, middlewares.CheckLogin)
	book.GET("/", bc.bookCtr.GetBook)
	book.PUT("/:bookID", bc.bookCtr.UpdateBook, middlewares.CheckLogin)
	book.DELETE("/:bookID", bc.bookCtr.DeleteBook, middlewares.CheckLogin)
}

