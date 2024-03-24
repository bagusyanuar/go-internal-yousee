package middleware

import (
	"encoding/json"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

type SessionMiddleware struct {
	CookieAuthConfig *common.CookieAuthConfig
}

func NewSessionMiddleware(cookieAuthConfig *common.CookieAuthConfig) SessionMiddleware {
	return SessionMiddleware{
		CookieAuthConfig: cookieAuthConfig,
	}
}

func (c *SessionMiddleware) CookieAuth() fiber.Handler {
	return encryptcookie.New(encryptcookie.Config{
		Key: c.CookieAuthConfig.SecretKey,
	})
}

func (c *SessionMiddleware) Verify() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		val := ctx.Cookies(c.CookieAuthConfig.CookieName)
		// fmt.Printf("session value : %+v", val)
		if val == "" {
			return common.JSONUnauthorized(ctx, "unauthorized", nil)
		}
		var user entity.User
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {
			return common.JSONUnauthorized(ctx, "unauthorized", nil)
		}
		ctx.Locals("rbac", user)
		return ctx.Next()
	}
}
