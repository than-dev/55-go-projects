package models

import (
	"github.com/than-dev/55-go-projects/3-api+mysql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	// db.NewRecord(book)
	db.Create(&book)

	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	
	return Books
}

func GetBookById(id string) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(id string) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}