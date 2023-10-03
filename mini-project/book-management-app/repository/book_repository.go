package repository

import (
	"book-app/database"
	"book-app/model"
	"errors"
)

func FindAll() []model.Book {
	return database.BookList
}

func FindById(id int) (model.Book, error) {
	for _, book := range database.BookList {
		if book.ID == id {
			return book, nil
		}
	}
	return model.Book{}, errors.New("book not found")
}

func Save(book model.Book) model.Book {
	database.BookList = append(database.BookList, book)
	return book
}

func UpdateById(id int, book model.Book) (model.Book, error) {
	for index, oldBook := range database.BookList {
		if oldBook.ID == id {
			database.BookList[index] = book
			return book, nil
		}
	}
	return model.Book{}, errors.New("book not found")
}

func DeleteById(id int) error {
	for index, book := range database.BookList {
		if book.ID == id {
			database.BookList = append(database.BookList[:index], database.BookList[index+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
