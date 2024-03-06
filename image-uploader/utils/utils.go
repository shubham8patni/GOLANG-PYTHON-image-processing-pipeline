package utils

import (
	"github.com/google/uuid"
	"io"
	"os"
)

func GenerateUUID() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func StoreImageAtPath(path string, image io.Reader) error {
	storedFile, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(storedFile, image)
	if err != nil {
		return err
	}
	return nil
}
