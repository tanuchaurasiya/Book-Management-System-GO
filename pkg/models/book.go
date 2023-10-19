package models

import (
	"book_management_system_golang/pkg/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID=?", id).Find(&getBook)
	return &getBook, db

}

func DeleteBook(id int64) Book {
	var book Book
	fmt.Println("this is id=", id)
	db.Where("ID=?", id).Delete(book)
	fmt.Println("deletbook", book)
	return book 
}
