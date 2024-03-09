package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type AuthController struct {
	Config      *viper.Viper
	AuthService service.AuthService
}

func NewAuthController(config *viper.Viper, authService service.AuthService) *AuthController {
	return &AuthController{
		Config:      config,
		AuthService: authService,
	}
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {

	username, err := c.AuthService.SignIn(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.JSON(&fiber.Map{
		"code":    200,
		"message": "success",
		"data":    username,
	})
}
