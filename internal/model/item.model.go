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
