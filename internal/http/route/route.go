package route

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/http/controller"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	HomeController *controller.HomeController
	AuthController *controller.AuthController
}

func (c *RouteConfig) Setup() {
	c.GuestRoute()
}

func (c *RouteConfig) GuestRoute() {
	c.App.Get("/", c.HomeController.Index)
	c.App.Get("/sign-in", c.AuthController.SignIn)
}
