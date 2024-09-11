package Repository

import (
	"Book/Config"
	"Book/Model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
)

func GetBook(c *fiber.Ctx) error {
	db := Config.GetDB()

	var book []Model.Book

	//result := db.Find(&book).Preload("Author")
	result := db.Preload("Author").Find(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"error":      result.Error,
		})
	}
	return c.JSON(book)
}

func GetBookById(c *fiber.Ctx) error {
	db := Config.GetDB()
	var book []Model.Book
	result := db.Where("id = ?", c.Params("id")).Find(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"error":      result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Book successfully retrieved",
		"book":       book,
	})
}

func GetBookByAuthor(c *fiber.Ctx) error {
	db := Config.GetDB()
	var listAuthor []Model.Author
	var book []Model.Book
	var requestAuthor struct {
		Author *Model.Author `json:"author"`
	}

	log.Println(requestAuthor.Author)
	//log.Println(requestAuthor.Author.FirstName)
	//log.Println(requestAuthor.Author.LastName)

	if err := c.BodyParser(&requestAuthor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	if requestAuthor.Author != nil {
		if requestAuthor.Author.FirstName != "" && requestAuthor.Author.LastName != "" {
			err := db.Where("first_name = ? AND last_name = ?",
				requestAuthor.Author.FirstName,
				requestAuthor.Author.LastName).Find(&listAuthor).Error

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": fiber.StatusInternalServerError,
					"error":      err.Error(),
				})
			}
		} else if requestAuthor.Author.FirstName != "" {
			err := db.Where("LOWER(first_name) = LOWER(?)", requestAuthor.Author.FirstName).Find(&listAuthor).Error

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": fiber.StatusInternalServerError,
					"error":      err.Error(),
				})
			}
		} else if requestAuthor.Author.LastName != "" {
			err := db.Where("last_name = ?", requestAuthor.Author.LastName).Find(&listAuthor).Error
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": fiber.StatusInternalServerError,
					"error":      err.Error(),
				})
			}
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "No valid author information provided",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "No author information provided",
		})
	}

	authorID := make([]uuid.UUID, len(listAuthor))
	for i, author := range listAuthor {
		authorID[i] = author.Id
	}
	//return c.JSON(fiber.Map{
	//	"statusCode": fiber.StatusOK,
	//	"data":       listAuthor,
	//	"authorID":   authorID,
	//})

	errFindBook := db.Where("author_id IN ?", authorID).Preload("Author").Find(&book).Error
	if errFindBook != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"error":      errFindBook.Error,
		})
	}
	return c.JSON(book)
}

func GetBookByTitle(c *fiber.Ctx) error {
	db := Config.GetDB()
	var book []Model.Book
	result := db.Where("title = ?", c.Params("title")).Find(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"error":      result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Book successfully retrieved",
		"book":       book,
	})
}

func GetBookByPublisher(c *fiber.Ctx) error {
	db := Config.GetDB()
	var book []Model.Book
	result := db.Where("publisher= ?", c.Params("publisher")).Find(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"error":      result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Book successfully retrieved",
		"book":       book,
	})
}

func CreateBook(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["Role"].(string)

	if role != "Admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"statusCode": fiber.StatusForbidden,
			"message":    "Access denied, admin only",
		})
	}

	db := Config.GetDB()

	//var newbook Model.Book

	var requestDataBook struct {
		Title     string        `json:"title"`
		Author    *Model.Author `json:"author"`
		Publisher string        `json:"publisher"`
	}

	if err := c.BodyParser(&requestDataBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	var authorId uuid.UUID
	if requestDataBook.Author != nil {
		var existsAuthor Model.Author

		findAuthor := db.Where("first_name = ? AND last_name = ?",
			requestDataBook.Author.FirstName,
			requestDataBook.Author.LastName).
			First(&existsAuthor).Error

		if findAuthor == nil {
			authorId = existsAuthor.Id
		} else {
			newAuthor := Model.Author{
				Id:        uuid.New(),
				FirstName: requestDataBook.Author.FirstName,
				LastName:  requestDataBook.Author.LastName,
			}

			err := db.Create(&newAuthor).Error
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"statusCode": fiber.StatusInternalServerError,
					"error":      err.Error(),
				})
			}
			authorId = newAuthor.Id
		}
	}

	newBook := Model.Book{
		Title:     requestDataBook.Title,
		AuthorID:  authorId,
		Publisher: requestDataBook.Publisher,
	}
	newBook.Id = uuid.New()
	result := db.Create(&newBook)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"statusCode": fiber.StatusCreated,
		"message":    "Book created successfully",
	})
}
