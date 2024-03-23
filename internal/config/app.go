package config

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/http/controller"
	"github.com/bagusyanuar/go-internal-yousee/internal/http/middleware"
	"github.com/bagusyanuar/go-internal-yousee/internal/http/route"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App        *fiber.App
	DB         *gorm.DB
	Log        *logrus.Logger
	Config     *viper.Viper
	JWT        *common.JWT
	Validator  *validator.Validate
	CookieAuth *common.CookieAuthConfig
}

func Bootstrap(config *BootstrapConfig) {
	jwtMiddleware := middleware.NewJWTMiddleware(config.JWT)
	sessionMiddleware := middleware.NewSessionMiddleware(config.CookieAuth)

	authRepository := repositories.NewAuthRepository(config.DB, config.Log)
	dashboardRepository := repositories.NewDashboardRepository(config.DB, config.Log)
	typeRepository := repositories.NewTypeRepository(config.DB, config.Log)
	provinceRepository := repositories.NewProvinceRepository(config.DB, config.Log)
	cityRepository := repositories.NewCityRepository(config.DB, config.Log)
	vendorRepository := repositories.NewVendorRepository(config.DB, config.Log)
	itemRepository := repositories.NewItemRepository(config.DB, config.Log)

	authService := service.NewAuthService(authRepository, config.JWT, config.Validator)
	dashboardService := service.NewDashboardService(dashboardRepository, config.Log)
	typeService := service.NewItemTypeService(typeRepository, config.Log, config.Validator)
	provinceService := service.NewProvinceService(provinceRepository, config.Log)
	cityService := service.NewCityService(cityRepository, config.Log)
	vendorService := service.NewVendorService(vendorRepository, config.Log, config.Validator)
	itemService := service.NewItemService(itemRepository, config.Log, config.Validator)

	homeController := controller.NewHomeController(config.Config)
	authController := controller.NewAuthController(config.Config, authService, config.Log, config.CookieAuth)
	dashboardController := controller.NewDashboardController(dashboardService, config.Log)
	typeController := controller.NewTypeController(typeService, config.Log)
	provinceController := controller.NewProvinceController(provinceService, config.Log)
	cityController := controller.NewCityController(cityService, config.Log)
	vendorController := controller.NewVendorController(vendorService, config.Log)
	itemController := controller.NewItemController(itemService, config.Log)

	routeConfig := route.RouteConfig{
		App:                 config.App,
		JWTMiddleware:       &jwtMiddleware,
		SessionMiddleware:   &sessionMiddleware,
		HomeController:      homeController,
		AuthController:      authController,
		DashboardController: dashboardController,
		TypeController:      typeController,
		ProvinceController:  provinceController,
		CityController:      cityController,
		VendorController:    vendorController,
		ItemController:      itemController,
	}
	routeConfig.Setup()
}
