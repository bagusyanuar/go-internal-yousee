package controller

import (
	"encoding/json"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthController struct {
	Config           *viper.Viper
	AuthService      service.AuthService
	Log              *logrus.Logger
	CookieAuthConfig *common.CookieAuthConfig
}

func NewAuthController(config *viper.Viper, authService service.AuthService, log *logrus.Logger, cookieAuthConfig *common.CookieAuthConfig) *AuthController {
	return &AuthController{
		Config:           config,
		AuthService:      authService,
		Log:              log,
		CookieAuthConfig: cookieAuthConfig,
	}
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {

	request := new(model.AuthRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, "failed to parse request body", nil)
	}

	//validate form request
	validation := c.AuthService.ValidateFormRequest(ctx.UserContext(), request)
	if validation.Error != nil {
		return common.JSONBadRequest(ctx, "invalid form request", validation.Data)
	}

	response, user := c.AuthService.SignIn(ctx.UserContext(), request)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Error.Error())
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	cookieValue, err := json.Marshal(&user)
	if err != nil {
		c.Log.Warnf("failed to marshal : %+v", err.Error())
		return common.JSONFromError(ctx, common.StatusInternalServerError, err, nil)
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     c.CookieAuthConfig.CookieName,
		Value:    string(cookieValue),
		Path:     "/",
		MaxAge:   c.CookieAuthConfig.MaxAge,
		SameSite: "strict",
	})
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully login",
		Data:    response.Data,
	})
}
