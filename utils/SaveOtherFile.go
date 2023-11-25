package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/registrar/config"
)

func SaveOtherFile(filename, lastname, controlNumber, key string, file *multipart.FileHeader, c *gin.Context) (string, string, error) {
	transformedFilename := strings.ReplaceAll(filename, " ", "-")
	ext := filepath.Ext(file.Filename)
	newFilename := transformedFilename + ext
	path := config.BASE_FOLDER + newFilename

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return newFilename, path, err
	}

	return newFilename, path, nil
}
