package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleUpdateOtherFile(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	form, err := c.MultipartForm()
	if err != nil {
        logrus.Warn("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=invalid_form")
		return
	}

	filename := form.Value["filename"][0]
	location := form.Value["fileLocation"][0]

	for key := range form.File {
		file := form.File[key][0]

		// Only allow pdf File
		if utils.IsFilePdf(file) == false {
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=not_pdf")
			return
		}

		remoteLocation, err := utils.UpdateOtherFile(c, file, filename, location)
		if err != nil {
            logrus.Warn("err saving updated file: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=upload_failed")
			return
		}

		err = store.UpdateOtherFile(remoteLocation, filename)
		if err != nil {
            logrus.Warn("err updating file: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed&reason=update_failed")
			return
		}
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusFound, url+"?status=success")
}
