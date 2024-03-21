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

	authMiddleware := c.JWTMiddleware.Verify()
	c.ProtectedRoute(apiRoute, authMiddleware)
}

func (c *RouteConfig) PublicRoute(route fiber.Router) {
	route.Get("/", c.HomeController.Index)
	route.Post("/sign-in", c.AuthController.SignIn)

}

func (c *RouteConfig) ProtectedRoute(route fiber.Router, authMiddleware fiber.Handler) {

	//media type routes
	typeGroup := route.Group("/media-type", authMiddleware)
	typeGroup.Get("/", c.TypeController.FindAll)
	typeGroup.Post("/", c.TypeController.Create)
	typeGroup.Get("/:id", c.TypeController.FindByID)
	typeGroup.Put("/:id", c.TypeController.Patch)
	typeGroup.Delete("/:id", c.TypeController.Delete)

	provinceGroup := route.Group("/province", authMiddleware)
	provinceGroup.Get("/", c.ProvinceController.FindAll)
	provinceGroup.Get("/:id", c.ProvinceController.FindByID)

	cityGroup := route.Group("/city", authMiddleware)
	cityGroup.Get("/", c.CityController.FindAll)
	cityGroup.Get("/:id", c.CityController.FindByID)

	vendorGroup := route.Group("/vendor", authMiddleware)
	vendorGroup.Get("/", c.VendorController.FindAll)
	vendorGroup.Post("/", c.VendorController.Create)
	vendorGroup.Get("/:id", c.VendorController.FindByID)
	vendorGroup.Put("/:id", c.VendorController.Patch)
	vendorGroup.Delete("/:id", c.VendorController.Delete)

	itemGrop := route.Group("/item", authMiddleware)
	itemGrop.Get("/", c.ItemController.FindAll)
	itemGrop.Post("/", c.ItemController.Create)
	itemGrop.Get("/:id", c.ItemController.FindByID)
}
