package domain

import (
	"book-store/pkg/models"
	"book-store/pkg/types"
)

// IAuthorRepo defines the interface for database repository operations related to authors.
type IAuthorRepo interface {
	GetAuthor(authorID uint) ([]models.Author, error)
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(authorID uint) error
}

// IAuthorService defines the interface for service operations related to authors.
type IAuthorService interface {
	GetAuthor(authorID uint) ([]types.AuthorRequest, error)
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(authorID uint) error
}
