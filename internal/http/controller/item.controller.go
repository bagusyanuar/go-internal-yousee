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

type (
	ItemController struct {
		ItemService service.ItemService
		Log         *logrus.Logger
	}
)

func NewItemController(itemService service.ItemService, log *logrus.Logger) *ItemController {
	return &ItemController{
		ItemService: itemService,
		Log:         log,
	}
}

func (c *ItemController) FindAll(ctx *fiber.Ctx) error {
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
	response, err := c.ItemService.FindAll(ctx.UserContext(), queryString)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show items",
		Data:    response.Data,
		Meta:    response.Meta,
	})
}

func (c *ItemController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res, err := c.ItemService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return common.JSONNotFound(ctx, err.Error(), nil)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show item",
		Data:    res,
	})
}
