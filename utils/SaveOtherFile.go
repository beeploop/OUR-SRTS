package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/gin-gonic/gin"
)

func SaveOtherFile(filename, lastname, controlNumber, key string, file *multipart.FileHeader, c *gin.Context) (string, string, error) {
	transformedFilename := strings.ReplaceAll(filename, " ", "-")
	ext := filepath.Ext(file.Filename)
	newFilename := transformedFilename + ext
	path := config.Env.BaseFolder + newFilename
	location := "http://" + config.Env.LocalAddr + config.Env.Port + "/" + path

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return newFilename, location, err
	}

	return newFilename, location, nil
}
