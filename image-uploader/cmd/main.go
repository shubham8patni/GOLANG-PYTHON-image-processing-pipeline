package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"image-uploader/db"
	"image-uploader/handlers"
	"image-uploader/models"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	db.DB.AutoMigrate(&models.Order{})
	e.POST("image_processing/upload", handlers.UploadAndProcessImage)
	e.Logger.Fatal(e.Start(":8800"))
}
