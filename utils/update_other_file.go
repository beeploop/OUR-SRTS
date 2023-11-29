package utils

import (
	"mime/multipart"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/gin-gonic/gin"
)

func UpdateOtherFile(c *gin.Context, file *multipart.FileHeader, filename, location string) (string, error) {
	tempLocation := config.Env.TempDir + filename

	err := c.SaveUploadedFile(file, tempLocation)
	if err != nil {
		return "", err
	}

	remoteLocation, err := SaveFileToNas(tempLocation, "Other")
	if err != nil {
		return "", err
	}

	return remoteLocation, nil
}
