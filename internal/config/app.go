package config

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/http/route"
	"github.com/gofiber/fiber/v2"
)

type BootstrapConfig struct {
	App *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	routeConfig := route.RouteConfig{
		App: config.App,
	}
	routeConfig.Setup()
}
