package config

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		TrustedProxies: []string{"127.0.0.1", "localhost"},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if err != nil {
				if code == 404 {
					return c.Status(code).JSON(&fiber.Map{
						"code":    code,
						"message": "route not found",
						"data":    nil,
					})
				}
				return c.Status(code).JSON(&fiber.Map{
					"code":    code,
					"message": err.Error(),
					"data":    nil,
				})
			}
			return nil
		},
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin, Access-Control-Allow-Headers",
		AllowOrigins:     "http://127.0.0.1:3000",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Static("/assets", "./assets")
	return app
}
