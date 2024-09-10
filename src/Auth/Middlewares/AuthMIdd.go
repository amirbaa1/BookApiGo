package Middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

//func AuthMiddleware(secret string) fiber.Handler {
//	return jwtware.New(jwtware.Config{
//		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
//	})
//}

func AuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"StatusCode": fiber.StatusUnauthorized,
					"message":    "Missing or invalid JWT",
				})
			}
			return ctx.Next()
		},
	})
}
