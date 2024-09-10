package main

import (
	"Book/Config"
	"Book/Model"
	"Book/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func ConnectDatabase() {
	Config.Connect()
	log.Println("Connected to database")
	db := Config.GetDB()

	err := db.AutoMigrate(&Model.Book{}, &Model.Author{})
	if err != nil {
		if err != nil {
			log.Println("Failed to migrate database: %v", err)
		}
	}
	log.Println("Database migrated")
}

func main() {
	app := fiber.New()

	ConnectDatabase()

	app.Use(logger.New())

	Routes.SetUpRouteBook(app)
	Routes.SetUpRouteAuthor(app)

	log.Fatal(app.Listen(":3000"))

}
