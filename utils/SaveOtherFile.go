package utils

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/gin-gonic/gin"
)

/*
Returns filename,location, and an error.

filename = name of the file
location = url of the file
error = error
*/
func SaveOtherFile(filename, lastname, controlNumber, key string, file *multipart.FileHeader, c *gin.Context) (string, string, error) {
	transformedFilename := strings.ReplaceAll(filename, " ", "-")
	ext := filepath.Ext(file.Filename)
	newFilename := lastname + "_" + controlNumber + "_Other_" + transformedFilename + ext
	path := config.Env.TempDir+ newFilename
	// location := "http://" + config.Env.LocalAddr + config.Env.Port + "/" + path

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return newFilename, "", err
	}

	remoteFileLocation, err := SaveFileToNas(path, key)
	if err != nil {
		return "", path, err
	}

	// delete the file in the temp dir
	err = os.Remove(path)
	if err != nil {
		return "", path, err
	}

	return newFilename, remoteFileLocation, nil
}
