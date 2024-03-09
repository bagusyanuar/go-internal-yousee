package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type HomeController struct {
	Config *viper.Viper
}

func NewHomeController(config *viper.Viper) *HomeController {
	return &HomeController{
		Config: config,
	}
}

func (c *HomeController) Index(ctx *fiber.Ctx) error {
	appName := c.Config.GetString("APP_NAME")
	appVersion := c.Config.GetString("APP_VERSION")
	return ctx.JSON(&fiber.Map{
		"app_name": appName,
		"version":  appVersion,
	})
}
