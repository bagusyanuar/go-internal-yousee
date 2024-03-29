package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TypeController struct {
	TypeService service.TypeService
	Log         *logrus.Logger
}

func NewTypeController(typeService service.TypeService, log *logrus.Logger) *TypeController {
	return &TypeController{
		TypeService: typeService,
		Log:         log,
	}
}

func (c *TypeController) FindAll(ctx *fiber.Ctx) error {
	rbac := ctx.Locals("rbac")
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

	c.Log.Warnf("user : %+v", rbac)
	response := c.TypeService.FindAll(ctx.UserContext(), queryString)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show media types",
		Data:    response.Data,
		Meta:    response.MetaPagination,
	})
}

func (c *TypeController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.TypeService.FindByID(ctx.UserContext(), id)
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show type",
		Data:    response.Data,
	})
}

func (c *TypeController) Create(ctx *fiber.Ctx) error {
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, "failed to parse request body", nil)
	}

	//validate form request
	validation := c.TypeService.ValidateFormRequest(ctx.UserContext(), request)
	if validation.Error != nil {
		return common.JSONBadRequest(ctx, "invalid form request", validation.Data)
	}

	//parsing multipart file
	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["icon"]
		for _, file := range files {
			request.Icon = file
		}
	}

	response := c.TypeService.Create(ctx.UserContext(), request)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Error.Error())
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONCreated(ctx, common.ResponseMap{
		Message: "successfully create media type",
	})
}

func (c *TypeController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, "failed to parse request body", nil)
	}

	//validate form request
	validation := c.TypeService.ValidateFormRequest(ctx.UserContext(), request)
	if validation.Error != nil {
		c.Log.Warnf("invalid form request : %+v", err)
		return common.JSONBadRequest(ctx, "invalid form request", validation.Data)
	}

	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["icon"]
		for _, file := range files {
			request.Icon = file
		}
	}

	response := c.TypeService.Patch(ctx.UserContext(), id, request)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Error.Error())
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully update media type",
	})
}

func (c *TypeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response := c.TypeService.Delete(ctx.UserContext(), id)
	if response.Error != nil {
		c.Log.Warnf("failed : %+v", response.Error.Error())
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully delete media type",
	})
}
