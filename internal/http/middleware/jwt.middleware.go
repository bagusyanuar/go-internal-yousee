package middleware

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type JWTMiddleware struct {
	Config *common.JWT
	Log    *logrus.Logger
}

func NewJWTMiddleware(config *common.JWT, log *logrus.Logger) JWTMiddleware {
	return JWTMiddleware{
		Config: config,
	}
}

func (j *JWTMiddleware) Verify() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			j.Log.Warnf("header auth : %+v", c.Get("Authorization"))
			return common.JSONUnauthorized(c, err.Error(), nil)
		},
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(j.Config.SignatureKey),
		},
	})
}
