package controllers

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"
	"book-store/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// to access the methods of service and repo
type BookController struct {
	bookSvc domain.IBookService
}

func NewBookController(bookSvc domain.IBookService) BookController {
	return BookController{
		bookSvc: bookSvc,
	}
}

func (bs *BookController) CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	book := &models.Book{
		BookName:    reqBook.BookName,
		AuthorId:      reqBook.AuthorId,
		Description: reqBook.Description,
	}

	if err := bs.bookSvc.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Book was created successfully")
}
func (bs *BookController) GetBook(e echo.Context) error {
	tempBookID := e.QueryParam("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil && tempBookID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	book, err := bs.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, book)
}

func (bs *BookController) UpdateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	existingBook, err := bs.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook := &models.Book{
		ID:          uint(bookID),
		BookName:    reqBook.BookName,
		AuthorId: reqBook.AuthorId,
		Description: reqBook.Description,
	}
	if updatedBook.BookName == "" {
		updatedBook.BookName = existingBook[0].BookName
	}
	if updatedBook.AuthorId == 0 {
		updatedBook.AuthorId = existingBook[0].AuthorId
	}
	if updatedBook.Description == "" {
		updatedBook.Description = existingBook[0].Description
	}
	if err := bs.bookSvc.UpdateBook(updatedBook); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Book was updated successfully")
}

func (bs *BookController) DeleteBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = bs.bookSvc.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := bs.bookSvc.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Book was deleted successfully")
}
