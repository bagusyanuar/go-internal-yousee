package common

import (
	"github.com/gofiber/fiber/v2"
)

const (
	StatusOK                   int = 200
	StatusBadRequest           int = 400
	StatusUnauthorized         int = 401
	StatusForbidden            int = 403
	StatusNotFound             int = 404
	StatusUnProccessableEntity int = 422
	StatusInternalServerError  int = 500
)

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
	Meta    any    `json:"meta"`
}

type ResponseMap struct {
	Message string
	Data    any
	Meta    any
}

func JSONSuccess(ctx *fiber.Ctx, mapResponse ResponseMap) error {
	status := 200
	message := "success"
	if mapResponse.Message != "" {
		message = mapResponse.Message
	}
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    mapResponse.Data,
		Message: message,
		Code:    status,
		Meta:    mapResponse.Meta,
	})
}

func JSONError(ctx *fiber.Ctx, message string, data any) error {
	status := 500
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    data,
		Message: message,
		Code:    status,
	})
}

func JSONBadRequest(ctx *fiber.Ctx, message string, data any) error {
	status := 400
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    data,
		Message: message,
		Code:    status,
	})
}

func JSONUnauthorized(ctx *fiber.Ctx, message string, data any) error {
	status := 401
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    data,
		Message: message,
		Code:    status,
	})
}

func JSONNotFound(ctx *fiber.Ctx, message string, data any) error {
	status := 404
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    data,
		Message: message,
		Code:    status,
	})
}

func JSONFromError(ctx *fiber.Ctx, code int, err error, data any) error {
	if code != 500 {
		return ctx.Status(code).JSON(APIResponse[any]{
			Data:    data,
			Message: err.Error(),
			Code:    code,
		})
	}
	return ctx.Status(500).JSON(APIResponse[any]{
		Data:    data,
		Message: err.Error(),
		Code:    500,
	})
}
