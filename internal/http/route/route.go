package route

import "github.com/gofiber/fiber/v2"

type RouteConfig struct {
	App *fiber.App
}

func (c *RouteConfig) Setup() {
	c.GuestRoute()
}

func (c *RouteConfig) GuestRoute() {
	c.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"app_name": "internal-yousee",
			"version":  "1.0.1",
		})
	})
}
