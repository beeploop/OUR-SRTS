package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const BASE_FOLDER = "documents/"

func FileSaver(c *gin.Context, file *multipart.FileHeader, lastname, controlNumber, filetype string) (string, error) {

	if StringInSlice(filetype) {
		fmt.Println("multiple entries")
		return "", errors.New("multiple entries")
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%v_%v_%v%v", lastname, controlNumber, filetype, ext)
	path := BASE_FOLDER + filename

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
