package Routes

import (
	"Auth/Middlewares"
	"Book/Repository"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouteBook(app *fiber.App) {
	jwtMiddleware := Middlewares.AuthMiddleware("SECRET:)")

	app.Get("/books", Repository.GetBook)
	app.Post("/books", jwtMiddleware, Repository.CreateBook)
	app.Get("/books/search", Repository.GetBookByAuthor)
	app.Get("/books/:id", Repository.GetBookById)
	app.Get("/books/:title", Repository.GetBookByTitle)
	app.Get("/books/:publisher", Repository.GetBookByPublisher)
}

func SetUpRouteAuthor(app *fiber.App) {
	app.Get("/authors", Repository.GetAuthorList)
}
