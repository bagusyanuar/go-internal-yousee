package model

import "mime/multipart"

type TypeRequest struct {
	Name string                `json:"name" validate:"required"`
	Icon *multipart.FileHeader `json:"icon"`
}

type TypeResponse struct {
	ID   uint64  `json:"id"`
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}
