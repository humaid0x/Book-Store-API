package domain

import (
	"book-store/pkg/models"
	"book-store/pkg/types"
)

// for database repository operation (call from service)
type IBookRepo interface {
	GetBooks(bookID uint) []models.Book
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}

// for service operation (response to controller | call from controller)
type IBookService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}
