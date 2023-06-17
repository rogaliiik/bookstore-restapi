package models

import (
	"github.com/jinzhu/gorm"

	"github.com/rogaliiik/bookstore/pkg/config"
)

var db *gorm.DB

// Model of book in database
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize of DB and migrate table
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Allows create new instance of book in DB
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Call all the books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Call specific book by Id
func GetBookById(Id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// Delete specific book by Id
func DeleteBook(Id int) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
