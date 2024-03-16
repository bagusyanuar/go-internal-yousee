package middleware

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authorization := ctx.Get("Authorization")
	if authorization == "" {

		return common.JSONUnauthorized(ctx, "unauthorized", nil)
	}
	getTokenClaim(authorization)
	return ctx.Next()
}

func getTokenClaim(header string) {

}
