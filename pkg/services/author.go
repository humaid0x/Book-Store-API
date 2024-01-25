package services

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"
	"book-store/pkg/types"
	"errors"
)

// authorService is the struct implementing IAuthorService
type authorService struct {
	repo domain.IAuthorRepo
}

// AuthorServiceInstance creates an instance of IAuthorService
func AuthorServiceInstance(authorRepo domain.IAuthorRepo) domain.IAuthorService {
	return &authorService{
		repo: authorRepo,
	}
}

// GetAuthor gets author information by ID
func (service *authorService) GetAuthor(authorID uint) ([]types.AuthorRequest, error) {
	var allAuthors []types.AuthorRequest
	authors, err := service.repo.GetAuthor(authorID)
	if err != nil {
		return nil, errors.New("error getting author")
	}
	if len(authors) == 0 {
		return nil, errors.New("no author found")
	}
	for _, val := range authors {
		allAuthors = append(allAuthors, types.AuthorRequest{
			ID:          val.ID,
			Name:        val.Name,
			Address:     val.Address,
			PhoneNumber: val.PhoneNumber,
		})
	}
	return allAuthors, nil
}

// CreateAuthor creates a new author
func (service *authorService) CreateAuthor(author *models.Author) error {
	if err := service.repo.CreateAuthor(author); err != nil {
		return errors.New("author was not created")
	}
	return nil
}

// UpdateAuthor updates an existing author
func (service *authorService) UpdateAuthor(author *models.Author) error {
	if err := service.repo.UpdateAuthor(author); err != nil {
		return errors.New("author update was unsuccessful")
	}
	return nil
}

// DeleteAuthor deletes an author by ID
func (service *authorService) DeleteAuthor(authorID uint) error {
	if err := service.repo.DeleteAuthor(authorID); err != nil {
		return errors.New("author deletion was unsuccessful")
	}
	return nil
}
