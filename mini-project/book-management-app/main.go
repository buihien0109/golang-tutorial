package main

import (
	"book-app/router"
	"github.com/kataras/iris/v12"
)

func main() {
	println("Running...")

	// Create a new instance of the application.
	app := iris.New()

	// Load the template files.
	tmpl := iris.HTML("./resources/views", ".html")
	app.RegisterView(tmpl)

	// Serve templates with "tmpl" as the primary one.
	app.HandleDir("/assets", iris.Dir("./resources/static"))

	// register routes
	router.SetupRouter(app)

	app.Listen(":8888")
}
