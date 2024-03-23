package common

import "errors"

var (
	ErrBadRequest          = errors.New("invalid bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidBearer       = errors.New("invalid bearer type")
	ErrInvalidJWTParse     = errors.New("invalid jwt parse")
	ErrUserNotFound        = errors.New("user not found")
	ErrRecordNotFound      = errors.New("record not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnknown             = errors.New("unknown error")
	ErrValidateRequest     = errors.New("unprocessable request validation")
	ErrGenerateToken       = errors.New("failed to generate access token")
)
