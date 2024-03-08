package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	username := "root"
	password := ""
	host := "localhost"
	port := 3306
	database := "db_yousee"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
