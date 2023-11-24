package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FileSaver(c *gin.Context, file *multipart.FileHeader, lastname, controlNumber, filetype string) (string, error) {
	ext := filepath.Ext(file.Filename)
	baseFolder := "documents/"
	filename := fmt.Sprintf("%v_%v_%v%v", lastname, controlNumber, filetype, ext)
	path := baseFolder + filename

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
