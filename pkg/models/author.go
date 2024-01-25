package models

type Author struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string 
	Address     string 
	PhoneNumber string 
}
