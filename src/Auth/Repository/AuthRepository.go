package Repository

import (
	"Auth/Config"
	"Auth/Model"
	"Auth/Utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	db := Config.GetDB()
	var user []Model.Auth

	var requestRegister struct {
		UserName string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BodyParser(&requestRegister); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}
	err := db.Where("username = ?", requestRegister.UserName).First(&user).Error
	fmt.Println(err)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"message":    "username already created",
			"error":      err,
		})
	}

	passHash, err := Utils.GeneratorPassword(requestRegister.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"message":    "password failed",
			"error":      err.Error(),
		})
	}
	newUser := Model.Auth{
		Id:       uuid.New(),
		UserName: requestRegister.UserName,
		Email:    requestRegister.Email,
		Password: passHash,
		Role:     "User",
	}

	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Create User Success",
	})

}

func Login(c *fiber.Ctx) error {
	db := Config.GetDB()

	var requestRegister struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user Model.Auth

	if err := c.BodyParser(&requestRegister); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	err := db.Where("user_name = ?", requestRegister.UserName).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}

	errPass := Utils.ValidatePassword(requestRegister.Password, user.Password)

	if errPass != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      errPass.Error(),
		})
	}

	claimsUser := Model.Auth{
		Id:       user.Id,
		UserName: user.UserName,
		Role:     user.Role,
	}

	token, err := Utils.GeneratorToken(claimsUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Login Success",
		"token":      token,
	})

}

func Profile(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token)
	claims := userId.Claims.(jwt.MapClaims)
	nameUser := claims["user_name"].(string)

	db := Config.GetDB()
	var userAuth Model.Auth

	err := db.Where("user_name = ?", nameUser).First(&userAuth).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"message":    "username not found",
			"error":      err.Error(),
		})
	}
	return c.SendString("Welcome " + userAuth.UserName + "\n" + userAuth.Email + "\n" + "Role: " + userAuth.Role)
}
