package rest

import (
	"book-app/model/request"
	"book-app/model/response"
	"book-app/service"
	"github.com/kataras/iris/v12"
	"net/http"
)

func GetAllBook(ctx iris.Context) {
	books := service.GetAllBooks()
	err := ctx.JSON(response.Success{
		Status:  http.StatusOK,
		Message: "get all books success",
		Data:    books,
	})
	if err != nil {
		return
	}
}

func GetBookById(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	book, err := service.GetBookById(id)
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	// return response success
	err = ctx.JSON(response.Success{
		Status:  http.StatusOK,
		Message: "get book by id success",
		Data:    book,
	})
}

func CreateBook(ctx iris.Context) {
	var bookRequest request.UpsertBookRequest
	err := ctx.ReadJSON(&bookRequest)
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	book := service.CreateBook(bookRequest)
	err = ctx.JSON(response.Success{
		Status:  http.StatusCreated,
		Message: "create book success",
		Data:    book,
	})
	if err != nil {
		return
	}
}

func UpdateBookById(ctx iris.Context) {
	var bookRequest request.UpsertBookRequest
	err := ctx.ReadJSON(&bookRequest)
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	book, err := service.UpdateBookById(id, bookRequest)
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = ctx.JSON(response.Success{
		Status:  http.StatusOK,
		Message: "update book success",
		Data:    book,
	})
	if err != nil {
		return
	}
}

func DeleteBookById(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	err = service.DeleteBookById(id)
	if err != nil {
		ctx.JSON(response.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	err = ctx.JSON(response.Success{
		Status:  http.StatusOK,
		Message: "delete book success",
		Data:    nil,
	})
	if err != nil {
		return
	}
}
