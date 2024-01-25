package services

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"
	"book-store/pkg/types"
	"errors"
)

// parent struct to implement interface binding
type bookService struct {
	repo domain.IBookRepo
}

// interface binding
func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &bookService{
		repo: bookRepo,
	}
}

// all methods of interface are implemented
func (service *bookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.repo.GetBooks(bookID)
	if len(book) == 0 {
		return nil, errors.New("no book found")
	}
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			ID:          val.ID,
			BookName:    val.BookName,
			AuthorId:      val.AuthorId,
			Description: val.Description,
		})
	}
	return allBooks, nil
}
func (service *bookService) CreateBook(book *models.Book) error {
	if err := service.repo.CreateBook(book); err != nil {
		return errors.New("book was not created")
	}
	return nil
}

func (service *bookService) UpdateBook(book *models.Book) error {
	if err := service.repo.UpdateBook(book); err != nil {
		return errors.New("book update was unsuccessful")
	}
	return nil
}
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.repo.DeleteBook(bookID); err != nil {
		return errors.New("book deletion was unsuccessful")
	}
	return nil
}
