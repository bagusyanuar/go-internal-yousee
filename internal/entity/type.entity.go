package entity

import "time"

type Type struct {
	ID        uint64    `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Icon      *string   `gorm:"column:icon" json:"username"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (t *Type) TableName() string {
	return "types"
}
