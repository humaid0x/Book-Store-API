package repositories

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"

	"gorm.io/gorm"
)

// authorRepo is the struct implementing IAuthorRepo
type authorRepo struct {
	db *gorm.DB
}

// AuthorDBInstance creates an instance of IAuthorRepo
func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}

// GetAuthor gets author information by ID
func (repo *authorRepo) GetAuthor(authorID uint) ([]models.Author, error) {
	var authors []models.Author
	var err error

	if authorID != 0 {
		err = repo.db.Where("id = ?", authorID).Find(&authors).Error
	} else {
		err = repo.db.Find(&authors).Error
	}
	if err != nil {
		return []models.Author{}, err
	}
	return authors, nil
}

// CreateAuthor creates a new author
func (repo *authorRepo) CreateAuthor(author *models.Author) error {
	if err := repo.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}

// UpdateAuthor updates an existing author
func (repo *authorRepo) UpdateAuthor(author *models.Author) error {
	if err := repo.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAuthor deletes an author by ID
func (repo *authorRepo) DeleteAuthor(authorID uint) error {
	var author models.Author
	if err := repo.db.Where("id = ?", authorID).First(&author).Error; err != nil {
		return err
	}

	// Delete books associated with the author
	if err := repo.db.Where("author_id = ?", authorID).Delete(&models.Book{}).Error; err != nil {
		return err
	}

	// Delete the author
	if err := repo.db.Delete(&author).Error; err != nil {
		return err
	}

	return nil
}

