package utils

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/sirupsen/logrus"
)

func SaveFileToNas(filePath, fileType string) (string, error) {
	nasUrl := config.Env.NasUrl + fileType + "/"

	err := os.MkdirAll(nasUrl, os.ModePerm)
	if err != nil {
		fmt.Println("failed to create directory: ", fileType)
		return "", err
	}

	// parse the nas url
	url, err := url.Parse(nasUrl)
	if err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, filename := filepath.Split(filePath)
	remoteFilename := filepath.Join(url.Path, filename)

	remoteFile, err := os.Create(remoteFilename)
	if err != nil {
		return remoteFilename, nil
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, file)
	if err != nil {
		return remoteFilename, nil
	}

	logrus.Info("File saved to NAS successfully")
	return remoteFilename, nil
}
