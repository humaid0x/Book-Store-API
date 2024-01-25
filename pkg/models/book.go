package models

type Book struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BookName    string
	AuthorId    uint
	Description string
}
