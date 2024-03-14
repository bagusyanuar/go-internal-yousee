package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	VendorController struct {
		VendorService service.VendorService
		Log           *logrus.Logger
	}
)

func NewVendorController(vendorService service.VendorService, log *logrus.Logger) *VendorController {
	return &VendorController{
		VendorService: vendorService,
		Log:           log,
	}
}

func (c *VendorController) FindAll(ctx *fiber.Ctx) error {
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
	response, err := c.VendorService.FindAll(ctx.UserContext(), queryString)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show vendors",
		Data:    response.Data,
		Meta:    response.Meta,
	})
}
