package route

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/http/controller"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	HomeController     *controller.HomeController
	AuthController     *controller.AuthController
	TypeController     *controller.TypeController
	ProvinceController *controller.ProvinceController
	CityController     *controller.CityController
}

func (c *RouteConfig) Setup() {
	c.GuestRoute()
	c.AuthRoute()
}

func (c *RouteConfig) GuestRoute() {
	c.App.Get("/", c.HomeController.Index)
	c.App.Post("/sign-in", c.AuthController.SignIn)

}

func (c *RouteConfig) AuthRoute() {

	//media type routes
	routeType := c.App.Group("/type")
	routeType.Get("/", c.TypeController.FindAll)
	routeType.Post("/", c.TypeController.Create)
	routeType.Get("/:id", c.TypeController.FindByID)
	routeType.Put("/:id", c.TypeController.Patch)
	routeType.Delete("/:id/delete", c.TypeController.Delete)

	routeProvince := c.App.Group("/province")
	routeProvince.Get("/", c.ProvinceController.FindAll)

	routeCity := c.App.Group("/city")
	routeCity.Get("/", c.CityController.FindAll)
	routeCity.Get("/:id", c.CityController.FindByID)
}
