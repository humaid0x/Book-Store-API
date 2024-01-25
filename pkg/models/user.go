package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email string 
	Password string 
	Phone string 
	Address string 
}