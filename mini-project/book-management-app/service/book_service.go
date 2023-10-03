package service

import (
	"book-app/model"
	"book-app/model/request"
	"book-app/repository"
	"github.com/google/uuid"
)

func GetAllBooks() []model.Book {
	return repository.FindAll()
}

func GetBookById(id int) (model.Book, error) {
	return repository.FindById(id)
}

func CreateBook(request request.UpsertBookRequest) model.Book {
	// generate id as number
	id := uuid.New().ID()
	book := model.Book{
		ID:          int(id),
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
	}
	return repository.Save(book)
}

func UpdateBookById(id int, request request.UpsertBookRequest) (model.Book, error) {
	bookUpdated := model.Book{
		ID:          id,
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
	}
	return repository.UpdateById(id, bookUpdated)
}

func DeleteBookById(id int) error {
	return repository.DeleteById(id)
}
