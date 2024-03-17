package config

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if err != nil {
				return c.Status(code).JSON(&fiber.Map{
					"code":    code,
					"message": err.Error(),
					"data":    nil,
				})
			}
			return nil
		},
	})
	return app
}
