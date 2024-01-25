package containers

import (
	"fmt"
	"log"

	"book-store/pkg/config"
	"book-store/pkg/connection"
	"book-store/pkg/controllers"
	"book-store/pkg/repositories"
	"book-store/pkg/routes"
	"book-store/pkg/services"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	config.SetConfig()
	db := connection.GetDB()

	// Book module setup
	bookRepo := repositories.BookDBInstance(db)
	bookService := services.BookServiceInstance(bookRepo)
	bookCtr := controllers.NewBookController(bookService)
	bookRoutes := routes.BookRoutes(e, bookCtr)
	bookRoutes.InitBookRoute()

	// Author module setup
	authorRepo := repositories.AuthorDBInstance(db)
	authorService := services.AuthorServiceInstance(authorRepo)
	authorCtr := controllers.NewAuthorController(authorService)
	authorRoutes := routes.AuthorRoutes(e, authorCtr)
	authorRoutes.InitAuthorRoute()

	// Auth module setup
	userRepo := repositories.UserDBInstance(db)
	authService := services.AuthServiceInstance(userRepo)
	authCtr := controllers.NewAuthController(authService)
	authRoutes := routes.AuthRoutes(e, authCtr)
	authRoutes.InitAuthRoute()


	// Starting the server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}

