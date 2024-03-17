package middleware

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	Config *common.JWT
}

func NewJWTMiddleware(config *common.JWT) JWTMiddleware {
	return JWTMiddleware{
		Config: config,
	}
}

func (j *JWTMiddleware) Verify() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return common.JSONUnauthorized(c, err.Error(), nil)
		},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(j.Config.SignatureKey),
		},
	})
}
