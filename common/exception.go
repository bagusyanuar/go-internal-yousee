package common

import "errors"

var (
	ErrBadRequest          = errors.New("invalid bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidBearer       = errors.New("invalid bearer type")
	ErrInvalidJWTParse     = errors.New("invalid jwt parse")
	ErrRecordNotFound      = errors.New("record not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
)
