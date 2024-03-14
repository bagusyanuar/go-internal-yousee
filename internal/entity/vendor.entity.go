package entity

import "time"

type Vendor struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"column:email;unique" json:"email"`
	Password  *string   `gorm:"column:password" json:"password"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Brand     string    `gorm:"column:brand" json:"brand"`
	PICName   string    `gorm:"column:picName" json:"pic_name"`
	PICPhone  string    `gorm:"column:picPhone" json:"pic_phone"`
	LastSeen  time.Time `gorm:"column:last_seen" json:"last_seen"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (v *Vendor) TableName() string {
	return "vendors"
}
