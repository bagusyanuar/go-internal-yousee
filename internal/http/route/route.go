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

	// apiRoute := c.App.Group("/api")
	apiRoute := c.App.Group("/api")
	c.PublicRoute(apiRoute)
	c.ProtectedRoute(apiRoute)
}

func (c *RouteConfig) PublicRoute(route fiber.Router) {
	route.Get("/", c.HomeController.Index)
	route.Post("/sign-in", c.AuthController.SignIn)

}

func (c *RouteConfig) ProtectedRoute(route fiber.Router) {

	//media type routes
	typeGroup := route.Group("/type", c.JWTMiddleware.Verify())
	typeGroup.Get("/", c.TypeController.FindAll)
	typeGroup.Post("/", c.TypeController.Create)
	typeGroup.Get("/:id", c.TypeController.FindByID)
	typeGroup.Put("/:id", c.TypeController.Patch)
	typeGroup.Delete("/:id/delete", c.TypeController.Delete)

	provinceGroup := route.Group("/province", c.JWTMiddleware.Verify())
	provinceGroup.Get("/", c.ProvinceController.FindAll)

	cityGroup := route.Group("/city", c.JWTMiddleware.Verify())
	cityGroup.Get("/", c.CityController.FindAll)
	cityGroup.Get("/:id", c.CityController.FindByID)

	vendorGroup := route.Group("/vendor", c.JWTMiddleware.Verify())
	vendorGroup.Get("/", c.VendorController.FindAll)
	vendorGroup.Post("/", c.VendorController.Create)
	vendorGroup.Get("/:id", c.VendorController.FindByID)
	vendorGroup.Put("/:id", c.VendorController.Patch)
	vendorGroup.Delete("/:id", c.VendorController.Delete)

	itemGrop := route.Group("/item", c.JWTMiddleware.Verify())
	itemGrop.Get("/", c.ItemController.FindAll)
	itemGrop.Get("/", c.ItemController.Create)
	itemGrop.Get("/:id", c.ItemController.FindByID)
}
