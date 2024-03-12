package controller

import (
	"errors"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CityController struct {
	CityService service.CityService
	Log         *logrus.Logger
}

func NewCityController(cityService service.CityService, log *logrus.Logger) *CityController {
	return &CityController{
		CityService: cityService,
		Log:         log,
	}
}

func (c *CityController) FindAll(ctx *fiber.Ctx) error {
	param := ctx.Query("name")
	page := ctx.QueryInt("page")
	perPage := ctx.QueryInt("per_page")

	pagination := model.PaginationQuery{
		Page:    page,
		PerPage: perPage,
	}

	queryString := model.QueryString[string]{
		Query:           param,
		QueryPagination: pagination,
	}
	response, err := c.CityService.FindAll(ctx.UserContext(), queryString)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show cities",
		Data:    response.Data,
		Meta:    response.Meta,
	})
}

func (c *CityController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res, err := c.CityService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return common.JSONNotFound(ctx, err.Error(), nil)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show city",
		Data:    res,
	})
}
