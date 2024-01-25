package repositories

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"

	"gorm.io/gorm"
)

// parent struct to implement interface binding
type bookRepo struct {
	db *gorm.DB
}

// interface binding
func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

// all methods of interface are implemented
func (repo *bookRepo) GetBooks(bookID uint) []models.Book {
	var book []models.Book
	var err error

	if bookID != 0 {
		err = repo.db.Where("id = ?", bookID).Find(&book).Error
	} else {
		err = repo.db.Find(&book).Error
	}
	if err != nil {
		return []models.Book{}
	}
	return book
}
func (repo *bookRepo) CreateBook(book *models.Book) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) UpdateBook(book *models.Book) error {
	if err := repo.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *bookRepo) DeleteBook(bookID uint) error {
	var Book models.Book
	if err := repo.db.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
