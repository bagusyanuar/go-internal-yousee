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
	response, code, err := c.TypeService.FindAll(ctx.UserContext(), queryString)
	if err != nil {
		return common.JSONFromError(ctx, code, err, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show media types",
		Data:    response.Data,
		Meta:    response.Meta,
	})
}

func (c *TypeController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, code, err := c.TypeService.FindByID(ctx.UserContext(), id)
	if err != nil {
		return common.JSONFromError(ctx, code, err, nil)
	}
	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show type",
		Data:    response,
	})
}

func (c *TypeController) Create(ctx *fiber.Ctx) error {
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["icon"]
		for _, file := range files {
			request.Icon = file
		}
	}

	code, validationMessage, err := c.TypeService.Create(ctx.UserContext(), request)
	if err != nil {
		return common.JSONFromError(ctx, code, err, validationMessage)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully create media type",
	})
}

func (c *TypeController) Patch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := new(model.TypeRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return common.JSONBadRequest(ctx, err.Error(), nil)
	}

	if form, err := ctx.MultipartForm(); err == nil {
		files := form.File["icon"]
		for _, file := range files {
			request.Icon = file
		}
	}

	err = c.TypeService.Patch(ctx.UserContext(), id, request)
	if err != nil {
		c.Log.Errorf("failed to patch data : %+v", err)
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully patch media type",
	})
}

func (c *TypeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.TypeService.Delete(ctx.UserContext(), id)
	if err != nil {
		return common.JSONError(ctx, err.Error(), nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully delete media type",
	})
}
