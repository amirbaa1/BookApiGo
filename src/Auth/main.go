package main

import (
	"Auth/Config"
	"Auth/Model"
	"Auth/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func ConnectDatabase() {
	Config.Connect()
	log.Println("Connected to database")
	db := Config.GetDB()

	err := db.AutoMigrate(&Model.Auth{})
	if err != nil {
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}
	log.Println("Database migrated")
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World TEST Fiber IdentityAPI!")
	})
	Routes.SetUpRoute(app)

	log.Fatal(app.Listen(":3001"))

}
