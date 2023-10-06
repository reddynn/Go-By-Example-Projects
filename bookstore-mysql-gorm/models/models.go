package models

import (
	"github.com/Go-By-Example-Projects/bookstore-mysql-gorm/config"
	"gorm.io/gorm"
)

type Book struct {
	Name        string `gorm:"varchar(255) " json:"name"`
	Author      string `gorm:"varchar(255)" json:"author"`
	Publication string `gorm:"varchar(255)" json:"publication"`
	ID          int    `gorm:"int;primaryKey;autoIncrement:true" json:"id,omitempty"`
}

var db *gorm.DB

func init() {
	db = config.Connect()
	db.AutoMigrate(&Book{})
}

func GetAll() []Book {
	var books []Book
	db.Find(&books)
	return books

}
func (b *Book) CreateBook() *Book {

	db.Create(&b)
	return b
}

func GetBookById(id int64) (*gorm.DB, *Book) {
	var getBook Book
	dbs := db.Where("id=?", id).Find(&getBook)
	return dbs, &getBook

}
