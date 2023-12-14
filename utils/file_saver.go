package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/gin-gonic/gin"
)

/*
Returns location, diskpath, and an error.

location = url of the file
diskpath = path of the file in the disk
error = error
*/
func FileSaver(c *gin.Context, file *multipart.FileHeader, lastname, controlNumber, filetype string) (string, string, error) {
	ext := filepath.Ext(file.Filename)

	// check if file is a multi-entry file
	if StringInSlice(filetype) {
		count, err := store.CountFilesInMultiEntry(filetype, controlNumber)
		if err != nil {
			fmt.Println("err counting: ", err)
			return "", "", err
		}

		strCount := strconv.Itoa(count)
		filename := lastname + "_" + controlNumber + "_" + filetype + "_" + strCount + ext
		path := filepath.Join(config.Env.TempDir, filename)
		// location := "http://" + config.Env.LocalAddr + config.Env.Port + "/" + path

		err = c.SaveUploadedFile(file, path)
		if err != nil {
			return "", path, err
		}

		nasFilepath, err := SaveFileToNas(path, filetype)
		if err != nil {
			return "", path, err
		}

		// delete the file in the temp dir
		err = os.Remove(path)
		if err != nil {
			return "", path, err
		}

		return nasFilepath, path, nil
	}

	filename := lastname + "_" + controlNumber + "_" + filetype + ext
	path := filepath.Join(config.Env.TempDir, filename)

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return "", path, err
	}

	nasFilepath, err := SaveFileToNas(path, filetype)
	if err != nil {
		return "", path, err
	}

	// delete the file in the temp dir
	err = os.Remove(path)
	if err != nil {
		return "", path, err
	}

	return nasFilepath, path, nil
}
