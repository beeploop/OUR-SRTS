package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
)

func HandleUpload(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed")
		return
	}

	lastname := form.Value["lastname"][0]
	controlNumber := form.Value["controlNumber"][0]

	// Use loop for getting file even though it will always be one file
	// because I'm too lazy to check what file is being uploaded
	for key := range form.File {
		file := form.File[key][0]

		if key != "Photo" {
			// Only allow pdf files
			if utils.IsFilePdf(file) == false {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed?reason=not_pdf")
				return
			}
		}

		// Separate method for handling Other files
		if key == "Other" {
			filename := form.Value["filename"][0]
			file, location, err := utils.SaveOtherFile(filename, lastname, controlNumber, key, file, c)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed")
				return
			}
			err = store.SaveOtherFile(file, location, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed")
				return
			}

		} else {
			location, err := utils.FileSaver(c, file, lastname, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed")
				return
			}
			err = store.SaveFile(location, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, url+"?status=failed")
				return
			}
		}
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
	return
}
