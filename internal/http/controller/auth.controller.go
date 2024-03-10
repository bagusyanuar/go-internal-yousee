package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthController struct {
	Config      *viper.Viper
	AuthService service.AuthService
	Log         *logrus.Logger
}

func NewAuthController(config *viper.Viper, authService service.AuthService, log *logrus.Logger) *AuthController {
	return &AuthController{
		Config:      config,
		AuthService: authService,
		Log:         log,
	}
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {

	request := new(model.AuthRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}
	res, err := c.AuthService.SignIn(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(common.APIResponse[*model.AuthResponse]{
		Data:    res,
		Message: "successfully login",
		Code:    200,
	})
}
