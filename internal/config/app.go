package config

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/http/route"
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
}

func Bootstrap(config *BootstrapConfig) {
	routeConfig := route.RouteConfig{
		App: config.App,
	}
	routeConfig.Setup()
}
