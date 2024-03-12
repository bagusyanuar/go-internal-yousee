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
	response, err := c.ProvinceService.FindAll(ctx.UserContext(), queryString)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show provinces",
		Data:    response.Data,
		Meta:    response.Meta,
	})
}
