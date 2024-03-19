package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProvinceController struct {
	ProvinceService service.ProvinceService
	Log             *logrus.Logger
}

func NewProvinceController(provinceService service.ProvinceService, log *logrus.Logger) *ProvinceController {
	return &ProvinceController{
		ProvinceService: provinceService,
		Log:             log,
	}
}

func (c *ProvinceController) FindAll(ctx *fiber.Ctx) error {
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
	response := c.ProvinceService.FindAll(ctx.UserContext(), queryString)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show provinces",
		Data:    response.Data,
		Meta:    response.MetaPagination,
	})
}

func (c *ProvinceController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.ProvinceService.FindByID(ctx.UserContext(), id)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show province",
		Data:    response.Data,
	})
}
