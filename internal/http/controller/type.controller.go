package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TypeController struct {
	TypeService service.TypeService
	Log         *logrus.Logger
}

func NewTypeController(typeService service.TypeService, log *logrus.Logger) *TypeController {
	return &TypeController{
		TypeService: typeService,
		Log:         log,
	}
}

func (c *TypeController) FindAll(ctx *fiber.Ctx) error {
	res, err := c.TypeService.FindAll(ctx.UserContext())
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(common.APIResponse[[]model.TypeResponse]{
		Data:    res,
		Message: "showing all data types",
		Code:    200,
	})
}

func (c *TypeController) Create(ctx *fiber.Ctx) error {
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}
	err = c.TypeService.Create(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(common.APIResponse[any]{
		Data:    nil,
		Message: "showing all data types",
		Code:    200,
	})
}
