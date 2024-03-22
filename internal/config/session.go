package config

import "github.com/gofiber/fiber/v2/middleware/session"

func NewAuthSession() *session.Store {
	config := session.Config{
		CookiePath: "/",
		KeyLookup:  "cookie:authentication-session",
	}
	store := session.New(config)
	return store
}
