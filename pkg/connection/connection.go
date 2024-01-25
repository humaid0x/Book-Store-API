package connection

import (
	"fmt"

	"book-store/pkg/config"
	"book-store/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// connect database
func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBHost, dbConfig.DBName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	db = d
}

// creating new table in bookstore database
func migrate() {
	db.Migrator().AutoMigrate(&models.Book{})
	db.Migrator().AutoMigrate(&models.Author{})
	db.Migrator().AutoMigrate(&models.User{})
}

// calling the connect function to initialize connection
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
