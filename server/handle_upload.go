package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/registrar/store"
	"github.com/registrar/utils"
)

func HandleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/search?status=failed")
		return
	}

	lastname := form.Value["lastname"][0]
	controlNumber := form.Value["controlNumber"][0]

	for key := range form.File {
		file := form.File[key][0]

		// Separate method for handling Other files
		if key == "Other" {
			filename := form.Value["filename"][0]
			file, location, err := utils.SaveOtherFile(filename, lastname, controlNumber, key, file, c)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, "/student/"+controlNumber+"?status=failed")
				return
			}
			err = store.SaveOtherFile(file, location, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, "/student/"+controlNumber+"?status=failed")
				return
			}

		} else {
			location, err := utils.FileSaver(c, file, lastname, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, "/student/"+controlNumber+"?status=failed")
				return
			}
			err = store.SaveFile(location, controlNumber, key)
			if err != nil {
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, "/student/"+controlNumber+"?status=failed")
				return
			}
		}

	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/student/"+controlNumber+"?status=success")
	return
}
