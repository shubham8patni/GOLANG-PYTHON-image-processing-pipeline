package handlers

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"image"
	"image-uploader/models"
	"image-uploader/utils"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"time"
)

func UploadAndProcessImage(c echo.Context) error {
	var email string
	form, err := c.MultipartForm() //c.FormValue("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if form.Value["email"] == nil {
		log.Println("email not received")
	} else {
		email = form.Value["email"][0]
	}

	uploadedImage := form.File["image"][0]
	if uploadedImage == nil {
		return c.JSON(http.StatusBadRequest, "server did not receive image")
	}

	src, err := uploadedImage.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not open image")
	}
	defer src.Close()

	// Read the file content into a buffer
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to read file content: ")
	}
	// Reset the file pointer before attempting to decode the configuration
	src.Seek(0, io.SeekStart)

	_, format, err := image.DecodeConfig(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to decode image")
	}

	switch format {
	case "jpeg", "png", "jpg":
		log.Println("image uploaded has valid format")
	default:
		log.Println("image uploaded has invalid format")
		return c.JSON(http.StatusBadRequest, "image should be in jpeg or png formats only")
	}

	uniqueID, err := utils.GenerateUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "something went wrong, can not generate unique key")
	}
	log.Println("unique ID generated : ", uniqueID)

	var imageName string
	if email == "" {
		imageName = string(time.DateTime) + uniqueID
	} else {
		imageName = string(time.DateTime) + email + uniqueID
	}
	log.Println("image name : ", imageName)

	storePath := utils.ImageStorePath + imageName + "." + format

	err = utils.StoreImageAtPath(storePath, buf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "server not able to store image")
	}

	modelEntry := &models.Order{
		PreProcessedImage: storePath,
		UserEmail:         email,
	}

	err = models.CreateOne(c, modelEntry)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to store data")
	}
	log.Println("data store in DB successfully")
	return c.JSON(http.StatusOK, "processing image")
}
