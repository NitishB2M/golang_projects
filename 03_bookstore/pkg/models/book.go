package models

import (
	"github.com/nitishb2m/golang_projects/03_bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	Author      string `gorm:"size:255;not null" json:"author"`
	Publication string `gorm:"size:255;" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) error {
	var book Book
	if err := db.Where("ID = ?", Id).First(&book).Error; err != nil {
		return err
	}
	if err := db.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}
