package Repository

import (
	"Book/Config"
	"Book/Model"
	"github.com/gofiber/fiber/v2"
)

func GetAuthorList(c *fiber.Ctx) error {
	db := Config.GetDB()
	var author []Model.Author

	err := db.Find(&author).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"error":      err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"authorList": author,
	})
}
