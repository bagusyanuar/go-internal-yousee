package model

import "time"

type VendorResponse struct {
	ID       uint64    `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Brand    string    `json:"brand"`
	PICName  string    `json:"pic_name"`
	PICPhone string    `json:"pic_phone"`
	LastSeen time.Time `json:"last_seen"`
}
