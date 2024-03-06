package models

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"image-uploader/db"
)

type Order struct {
	gorm.Model
	PreProcessedImage string `gorm:"column:pre_processed_image; required; unique"`
	UserEmail         string `gorm:"column:user_email"`
}

func CreateOne(c echo.Context, order *Order) error {
	err := db.DB.Create(order).Error
	return err
}
