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

	queryString := model.QueryString[string]{
		Query: param,
		QueryPagination: model.PaginationQuery{
			Page:    page,
			PerPage: perPage,
		},
	}

	response := c.VendorService.FindAll(ctx.UserContext(), queryString)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show vendors",
		Data:    response.Data,
		Meta:    response.MetaPagination,
	})
}

func (c *VendorController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.VendorService.FindByID(ctx.UserContext(), id)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show vendor",
		Data:    response.Data,
	})
}

func (c *VendorController) Create(ctx *fiber.Ctx) error {
	request := new(model.VendorRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, "failed to parse request body", nil)
	}

	//validate form request
	validation := c.VendorService.ValidateFormRequest(ctx.UserContext(), request)
	if validation.Error != nil {
		return common.JSONBadRequest(ctx, "invalid form request", validation.Data)
	}

	response := c.VendorService.Create(ctx.UserContext(), request)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Validation)
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONCreated(ctx, common.ResponseMap{
		Message: "successfully create vendor",
	})
}

func (c *VendorController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := new(model.VendorRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, "failed to parse request body", nil)
	}

	//validate form request
	validation := c.VendorService.ValidateFormRequest(ctx.UserContext(), request)
	if validation.Error != nil {
		c.Log.Warnf("invalid form request : %+v", err)
		return common.JSONBadRequest(ctx, "invalid form request", validation.Data)
	}

	response := c.VendorService.Patch(ctx.UserContext(), id, request)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Error.Error())
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully update vendor",
	})
}

func (c *VendorController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response := c.VendorService.Delete(ctx.UserContext(), id)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Error.Error())
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully delete vendor",
	})
}
