package model

import "time"

type ItemResponse struct {
	ID        uint64      `json:"id"`
	TypeID    *uint64     `json:"type_id"`
	CityID    *uint64     `json:"city_id"`
	VendorID  *uint64     `json:"vendor_id"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	Location  *string     `json:"location"`
	URL       *string     `json:"url"`
	Width     float64     `json:"width"`
	Height    float64     `json:"height"`
	Position  string      `json:"position"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	City      *ItemCity   `json:"city"`
	Type      *ItemType   `json:"type"`
	Vendor    *ItemVendor `json:"vendor"`
}

type ItemCity struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ItemType struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ItemVendor struct {
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

type ItemRequest struct {
	TypeID    *uint64 `json:"type_id,omitempty" validate:"required"`
	CityID    *uint64 `json:"city_id,omitempty" validate:"required"`
	VendorID  *uint64 `json:"vendor_id,omitempty" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Location  string  `json:"location" validate:"required"`
	URL       string  `json:"url" validate:"required"`
	Width     float64 `json:"width" validate:"required"`
	Height    float64 `json:"height" validate:"required"`
	Position  string  `json:"position" validate:"required"`
}

type ItemQueryString struct {
	Param    string `json:"param"`
	TypeID   string `json:"type_id"`
	CityID   string `json:"city_id"`
	VendorID string `json:"vendor_id"`
}
