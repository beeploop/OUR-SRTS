package utils

import (
	"mime/multipart"
	"path/filepath"
)

func IsFilePdf(file *multipart.FileHeader) bool {
	ext := filepath.Ext(file.Filename)
	return ext == ".pdf"
}
