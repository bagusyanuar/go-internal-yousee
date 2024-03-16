package model

import "time"

type VendorResponse struct {
	ID       uint64     `json:"id"`
	Email    string     `json:"email"`
	Name     string     `json:"name"`
	Address  string     `json:"address"`
	Phone    string     `json:"phone"`
	Brand    string     `json:"brand"`
	PICName  string     `json:"pic_name"`
	PICPhone string     `json:"pic_phone"`
	LastSeen *time.Time `json:"last_seen"`
}

type VendorRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Brand    string `json:"brand" validate:"required"`
	PICName  string `json:"pic_name"`
	PICPhone string `json:"pic_phone" validate:"required"`
}
