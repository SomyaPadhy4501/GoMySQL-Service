package models

import (
	"github.com/SomyaPadhy4501/book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Making book data

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialization

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//Creating the data base function
// Create Book function

func (b *Book) CreateABook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// All books

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

//Get Book by ID

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

//Delete book

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
