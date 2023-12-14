package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/sirupsen/logrus"
)

func SaveFileToNas(srcFilepath, fileType string) (string, error) {
	nasPath := filepath.Join(config.Env.NasUrl, fileType)

	err := os.MkdirAll(nasPath, os.ModePerm)
	if err != nil {
		fmt.Println("failed to create directory: ", fileType)
		return "", err
	}

	src, err := os.Open(srcFilepath)
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, filename := filepath.Split(srcFilepath)
	outputFilename := filepath.Join(nasPath, filename)

	distFile, err := os.Create(outputFilename)
	if err != nil {
		return outputFilename, nil
	}
	defer distFile.Close()

	_, err = io.Copy(distFile, src)
	if err != nil {
		return outputFilename, nil
	}

	logrus.Info("File saved to NAS successfully")
	return outputFilename, nil
}
