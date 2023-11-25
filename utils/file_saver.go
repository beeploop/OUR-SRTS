package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/registrar/store"
)

const BASE_FOLDER = "documents/"

func FileSaver(c *gin.Context, file *multipart.FileHeader, lastname, controlNumber, filetype string) (string, error) {
	ext := filepath.Ext(file.Filename)

	if StringInSlice(filetype) {
		fmt.Println("multiple entries")
		count, err := store.CountFilesInMultiEntry(filetype, controlNumber)
		if err != nil {
			fmt.Println("err counting: ", err)
		}
		fmt.Println("count: ", count)

		filename := fmt.Sprintf("%v_%v_%v_%v%v", lastname, controlNumber, filetype, count, ext)
		path := BASE_FOLDER + filename

		err = c.SaveUploadedFile(file, path)
		if err != nil {
			return path, err
		}

		return path, nil
	}

	filename := fmt.Sprintf("%v_%v_%v%v", lastname, controlNumber, filetype, ext)
	path := BASE_FOLDER + filename

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
