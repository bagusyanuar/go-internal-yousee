package entity

import "time"

type Item struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TypeID    *uint64   `gorm:"column:type_id" json:"type_id"`
	CityID    *uint64   `gorm:"column:city_id" json:"city_id"`
	VendorID  *uint64   `gorm:"column:vendor_id" json:"vendor_id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Latitude  float64   `gorm:"column:latitude" json:"latitude"`
	Longitude float64   `gorm:"column:longitude" json:"longitude"`
	Location  *string   `gorm:"column:location" json:"location"`
	URL       *string   `gorm:"column:url" json:"url"`
	Width     float64   `gorm:"column:width" json:"width"`
	Height    float64   `gorm:"column:height" json:"height"`
	Position  string    `gorm:"column:position" json:"position"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Type      *Type     `gorm:"foreignKey:TypeID" json:"type,omitempty"`
	City      *City     `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Vendor    *Vendor   `gorm:"foreignKey:VendorID" json:"vendor,omitempty"`
}

func (i *Item) TableName() string {
	return "items"
}
