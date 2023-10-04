package router

import (
	"book-app/controller"
	"book-app/rest"
	"github.com/kataras/iris/v12"
)

// SetupRouter set up router for application
func SetupRouter(app *iris.Application) {
	// View
	app.Get("/", controller.GetBookHomePage)
	app.Get("/books/create", controller.GetBookAddPage)
	app.Get("/books/{id}/detail", controller.GetBookDetailPage)

	// API
	v1 := app.Party("/api/v1/books")
	{
		v1.Get("/", rest.GetAllBook)
		v1.Get("/{id}", rest.GetBookById)
		v1.Post("/", rest.CreateBook)
		v1.Put("/{id}", rest.UpdateBookById)
		v1.Delete("/{id}", rest.DeleteBookById)
	}
}
