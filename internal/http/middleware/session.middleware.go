package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionMiddleware struct {
	CookieSession *session.Store
}

func NewSessionMiddleware(cookieSession *session.Store) SessionMiddleware {
	return SessionMiddleware{
		CookieSession: cookieSession,
	}
}

func (c *SessionMiddleware) Verify() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := c.CookieSession.Get(ctx)
		if err != nil {
			return ctx.Status(401).JSON(&fiber.Map{
				"code":    401,
				"message": err.Error(),
				"data":    nil,
			})
		}
		val := session.Get("authentication-session")
		fmt.Printf("session value : %+v", val)
		return ctx.Next()
	}
}
