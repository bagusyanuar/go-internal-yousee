package config

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/http/controller"
	"github.com/bagusyanuar/go-internal-yousee/internal/http/route"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App    *fiber.App
	DB     *gorm.DB
	Log    *logrus.Logger
	Config *viper.Viper
	JWT    *common.JWT
}

func Bootstrap(config *BootstrapConfig) {

	authRepository := repositories.NewAuthRepository(config.DB, config.Log)
	typeRepository := repositories.NewTypeRepository(config.DB, config.Log)

	authService := service.NewAuthService(authRepository, config.JWT)
	typeService := service.NewItemTypeService(typeRepository)

	homeController := controller.NewHomeController(config.Config)
	authController := controller.NewAuthController(config.Config, authService, config.Log)
	typeController := controller.NewTypeController(typeService, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		HomeController: homeController,
		AuthController: authController,
		TypeController: typeController,
	}
	routeConfig.Setup()
}
