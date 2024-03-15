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

func (c *VendorController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res, err := c.VendorService.FindByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return common.JSONNotFound(ctx, err.Error(), nil)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show vendor",
		Data:    res,
	})
}

func (c *VendorController) Create(ctx *fiber.Ctx) error {
	request := new(model.VendorRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	validationMsg, err := c.VendorService.Create(ctx.UserContext(), request)
	if err != nil {
		if errors.Is(common.ErrBadRequest, err) {
			return common.JSONBadRequest(ctx, "bad request", validationMsg)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully create vendor",
	})
}

func (c *VendorController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := new(model.VendorRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	validationMsg, err := c.VendorService.Patch(ctx.UserContext(), id, request)
	if err != nil {
		switch err {
		case common.ErrBadRequest:
			return common.JSONBadRequest(ctx, "bad request", validationMsg)
		case gorm.ErrRecordNotFound:
			return common.JSONNotFound(ctx, err.Error(), nil)
		default:
			return common.JSONError(ctx, err.Error(), nil)
		}
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully patch vendor",
	})
}

func (c *VendorController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.VendorService.Delete(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return common.JSONNotFound(ctx, err.Error(), nil)
		}
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully delete media type",
	})
}
