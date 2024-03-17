package middleware

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	// return func(ctx *fiber.Ctx) error {

	// 	authorization := ctx.Get("Authorization")
	// 	jwtClaim, err := j.verifyToken(authorization)
	// 	if err != nil {
	// 		return common.JSONUnauthorized(ctx, err.Error(), nil)
	// 	}
	// 	if jwtClaim != nil {
	// 		ctx.Set("user", "")
	// 	}

	// 	return ctx.Next()
	// 	return common.JSONUnauthorized(ctx, err.Error(), nil)
	// }
}

func (j *JWTMiddleware) verifyToken(header string) (*common.JWTClaims, error) {
	if header == "" {
		return nil, common.ErrUnauthorized
	}
	bearer := string(header[0:7])
	token := string(header[7:])
	if bearer != "Bearer " {
		return nil, common.ErrInvalidBearer
	}

	jwtSignatureKey := j.Config.SignatureKey
	t, e := jwt.ParseWithClaims(token, &common.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSignatureKey), nil
	})
	if e != nil {
		return nil, e
	}

	if claims, ok := t.Claims.(*common.JWTClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, common.ErrInvalidJWTParse
}
