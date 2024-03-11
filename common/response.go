package common

import "github.com/gofiber/fiber/v2"

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func JSONSuccess(ctx *fiber.Ctx, message string, data any) error {
	status := 200
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    data,
		Message: message,
		Code:    status,
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

func JSONNotFound(ctx *fiber.Ctx, message string, data any) error {
	status := 404
	return ctx.Status(status).JSON(APIResponse[any]{
		Data:    data,
		Message: message,
		Code:    status,
	})
}
