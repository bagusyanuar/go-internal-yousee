package common

import "errors"

var (
	ErrBadRequest    = errors.New("invalid bad request")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInvalidBearer = errors.New("invalid bearer type")
)
