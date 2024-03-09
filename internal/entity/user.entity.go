package entity

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`
	Email     string    `gorm:"column:email" json:"email"`
	Username  string    `gorm:"column:username" json:"username"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
