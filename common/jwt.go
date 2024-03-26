package common

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Issuer       string
	SignatureKey string
	Exp          uint
}

type JWTClaims struct {
	jwt.RegisteredClaims
	// UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
