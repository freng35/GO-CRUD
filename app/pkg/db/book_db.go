package database

import (
	"app/pkg/model"
	"gorm.io/gorm"
)

func CreateBook(book model.Book) {
	DB.Create(book)
}

func AddBook(book model.Book) {
	DB.Model(book).Where("name = ?", book.Name).Updates(
		map[string]interface{}{
			"amount": gorm.Expr("amount + ?", book.Amount)})
}

func GetBookByName(book model.Book) model.Book {
	var tmpBook model.Book
	DB.Find(&tmpBook, "name = ?", book.Name)
	return tmpBook
}

func GetBooks() []model.Book {
	var tmpBooks []model.Book
	DB.Find(&tmpBooks)
	return tmpBooks
}
