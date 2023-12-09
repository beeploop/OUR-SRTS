package server

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleUpdateFile(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	form, err := c.MultipartForm()
	if err != nil {
		logrus.Warn("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=invalid_form")
		return
	}

	controlNumber := form.Value["controlNumber"][0]

	// Use loop for getting file even though it will always be one file
	// because I'm too lazy to check what file is being uploaded
	for key := range form.File {
		file := form.File[key][0]

		if key != "Photo" {
			// Only allow pdf File
			if utils.IsFilePdf(file) == false {
				c.Request.Method = "GET"
				c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=not_pdf")
				return
			}
		}

		student, err := store.GetStudent(controlNumber)
		if err != nil {
			logrus.Warn("err getting student: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=unknown_student")
			return
		}

		location, _, err := utils.FileSaver(c, file, student.Lastname, controlNumber, key)
		if err != nil {
			logrus.Warn("err saving file: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=invalid_form")
			return
		}

		// replace backslash with forward slash for windows
		if runtime.GOOS == "windows" {
			location = strings.ReplaceAll(location, "\\", "/")
		}

		err = store.UpdateFile(key, controlNumber, location)
		if err != nil {
			logrus.Warn("err updating file: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=invalid_form")
			return
		}
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
