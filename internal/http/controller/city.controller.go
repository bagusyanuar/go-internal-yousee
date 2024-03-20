package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	province := ctx.Query("province")
	page := ctx.QueryInt("page")
	perPage := ctx.QueryInt("per_page")

	pagination := model.PaginationQuery{
		Page:    page,
		PerPage: perPage,
	}

	queryString := model.QueryString[model.CityQueryString]{
		Query: model.CityQueryString{
			Name:     param,
			Province: province,
		},
		QueryPagination: pagination,
	}
	response := c.CityService.FindAll(ctx.UserContext(), queryString)

	if response.Error != nil {
		c.Log.Warnf("service failed : %+v", response.Error)
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show cities",
		Data:    response.Data,
		Meta:    response.MetaPagination,
	})
}

func (c *CityController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.CityService.FindByID(ctx.UserContext(), id)

	if response.Error != nil {
		c.Log.Warnf("service failed : %+v", response.Error)
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show city",
		Data:    response.Data,
	})
}
