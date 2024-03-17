package route

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/http/controller"
	"github.com/bagusyanuar/go-internal-yousee/internal/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	JWTMiddleware      *middleware.JWTMiddleware
	HomeController     *controller.HomeController
	AuthController     *controller.AuthController
	TypeController     *controller.TypeController
	ProvinceController *controller.ProvinceController
	CityController     *controller.CityController
	VendorController   *controller.VendorController
	ItemController     *controller.ItemController
}

func (c *RouteConfig) Setup() {

	apiRoute := c.App.Group("/api")
	c.GuestRoute(apiRoute)
	apiRoute.Use(c.JWTMiddleware.Verify())
	c.ProtectedRoute(apiRoute)
}

func (c *RouteConfig) GuestRoute(apiRoute fiber.Router) {
	apiRoute.Get("/", c.HomeController.Index)
	apiRoute.Post("/sign-in", c.AuthController.SignIn)

}

func (c *RouteConfig) ProtectedRoute(apiRoute fiber.Router) {

	//media type routes
	typeGroup := apiRoute.Group("/type")
	typeGroup.Get("/", c.TypeController.FindAll)
	typeGroup.Post("/", c.TypeController.Create)
	typeGroup.Get("/:id", c.TypeController.FindByID)
	typeGroup.Put("/:id", c.TypeController.Patch)
	typeGroup.Delete("/:id/delete", c.TypeController.Delete)

	provinceGroup := apiRoute.Group("/province")
	provinceGroup.Get("/", c.ProvinceController.FindAll)

	cityGroup := apiRoute.Group("/city")
	cityGroup.Get("/", c.CityController.FindAll)
	cityGroup.Get("/:id", c.CityController.FindByID)

	vendorGroup := apiRoute.Group("/vendor")
	vendorGroup.Get("/", c.VendorController.FindAll)
	vendorGroup.Post("/", c.VendorController.Create)
	vendorGroup.Get("/:id", c.VendorController.FindByID)
	vendorGroup.Put("/:id", c.VendorController.Patch)
	vendorGroup.Delete("/:id", c.VendorController.Delete)

	itemGrop := apiRoute.Group("/item")
	itemGrop.Get("/", c.ItemController.FindAll)
	itemGrop.Get("/", c.ItemController.Create)
	itemGrop.Get("/:id", c.ItemController.FindByID)
}
