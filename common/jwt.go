package common

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Issuer       string
	SignatureKey string
}

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID uint64 `json:"user_Id"`
}
