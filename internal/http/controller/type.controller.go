package controller

import (
	"errors"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TypeController struct {
	TypeService service.TypeService
	Log         *logrus.Logger
	Validator   *validator.Validate
}

func NewTypeController(typeService service.TypeService, log *logrus.Logger, validator *validator.Validate) *TypeController {
	return &TypeController{
		TypeService: typeService,
		Log:         log,
		Validator:   validator,
	}
}

func (c *TypeController) FindAll(ctx *fiber.Ctx) error {
	param := ctx.Query("param")
	res, err := c.TypeService.FindAll(ctx.UserContext(), param)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show types",
		Data:    res,
	})
}

func (c *TypeController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res, err := c.TypeService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return common.JSONNotFound(ctx, err.Error(), nil)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show type",
		Data:    res,
	})
}

func (c *TypeController) Create(ctx *fiber.Ctx) error {
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	errValidation, msg := common.Validate(c.Validator, request)

	if errValidation != nil {
		c.Log.Warnf("validation value : %+v", msg)
		return common.JSONBadRequest(ctx, "bad request", msg)
	}

	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["icon"]
		for _, file := range files {
			request.Icon = file
		}
	}

	err = c.TypeService.Create(ctx.UserContext(), request)
	if err != nil {
		if errors.Is(common.ErrBadRequest, err) {
			return common.JSONBadRequest(ctx, err.Error(), nil)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully create media type",
	})
}

func (c *TypeController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["icon"]
		for _, file := range files {
			request.Icon = file
		}
	}

	err = c.TypeService.Patch(ctx.UserContext(), id, request)
	if err != nil {
		c.Log.Errorf("failed to patch data : %+v", err)
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully patch media type",
	})
}

func (c *TypeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.TypeService.Delete(ctx.UserContext(), id)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully delete media type",
	})
}
