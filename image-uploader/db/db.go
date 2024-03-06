package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "root:mysql@tcp(127.0.0.1:3309)/image_uploader?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	DB = db
	return nil
}
