package Routes

import (
	"Auth/Config"
	"Auth/Middlewares"
	"Auth/Repository"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(app *fiber.App) {
	app.Post("/auth/register", Repository.Register)
	app.Post("/auth/login", Repository.Login)

	jwtMiddleware := Middlewares.AuthMiddleware(string(Config.Secret))
	app.Get("/auth/profile", jwtMiddleware, Repository.Profile)
}
