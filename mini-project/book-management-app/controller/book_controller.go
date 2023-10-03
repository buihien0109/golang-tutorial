package controller

import (
	"book-app/database"
	"book-app/service"
	"github.com/kataras/iris/v12"
)

func GetBookHomePage(ctx iris.Context) {
	ctx.ViewData("BookList", database.BookList)
	err := ctx.View("index.html")
	if err != nil {
		return
	}
}

func GetBookAddPage(ctx iris.Context) {
	err := ctx.View("create.html")
	if err != nil {
		return
	}
}

func GetBookDetailPage(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		return
	}

	book, err := service.GetBookById(id)
	if err != nil {
		return
	}

	ctx.ViewData("Book", book)
	err = ctx.View("detail.html")
	if err != nil {
		return
	}
}
