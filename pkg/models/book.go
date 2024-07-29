package models

import (
	"firstwebproject/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) CreateBook() *Book {
	db.Model(b)
	db.Create(&b)
	return b
}
