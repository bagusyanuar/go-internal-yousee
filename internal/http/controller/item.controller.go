package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	param := ctx.Query("param")
	cityID := ctx.Query("city_id")
	typeID := ctx.Query("type_id")
	vendorID := ctx.Query("vendor_id")
	page := ctx.QueryInt("page")
	perPage := ctx.QueryInt("per_page")

	queryString := model.QueryString[model.ItemQueryString]{
		Query: model.ItemQueryString{
			Param:    param,
			CityID:   cityID,
			TypeID:   typeID,
			VendorID: vendorID,
		},
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

	response := c.ItemService.FindByID(ctx.UserContext(), id)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show item",
		Data:    response.Data,
	})
}

func (c *ItemController) Create(ctx *fiber.Ctx) error {
	request := new(model.ItemRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, "failed to parse request body", nil)
	}

	//validate form request
	validation := c.ItemService.ValidateFormRequest(ctx.UserContext(), request)
	if validation.Error != nil {
		return common.JSONBadRequest(ctx, "invalid form request", validation.Data)
	}

	response := c.ItemService.Create(ctx.UserContext(), request)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Validation)
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONCreated(ctx, common.ResponseMap{
		Message: "successfully create item",
	})
}
