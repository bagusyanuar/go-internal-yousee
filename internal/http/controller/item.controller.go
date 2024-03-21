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

	queryString := model.QueryString[string]{
		Query: param,
		QueryPagination: model.PaginationQuery{
			Page:    page,
			PerPage: perPage,
		},
	}

	response := c.ItemService.FindAll(ctx.UserContext(), queryString)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show vendors",
		Data:    response.Data,
		Meta:    response.MetaPagination,
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

func (c *ItemController) Create(ctx *fiber.Ctx) error {
	request := new(model.ItemRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	validationMsg, err := c.ItemService.Create(ctx.UserContext(), request)
	if err != nil {
		if errors.Is(common.ErrBadRequest, err) {
			return common.JSONBadRequest(ctx, "bad request", validationMsg)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully create new item",
	})
}
