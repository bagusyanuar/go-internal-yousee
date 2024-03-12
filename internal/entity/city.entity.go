package entity

import "time"

type City struct {
	ID         uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProvinceID uint64     `gorm:"column:province_id" json:"province_id"`
	Name       string     `gorm:"column:name" json:"name"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Province   *Province  `gorm:"foreignKey:ProvinceID" json:"province"`
}

func (c *City) TableName() string {
	return "cities"
}
