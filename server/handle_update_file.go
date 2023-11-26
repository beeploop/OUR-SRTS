package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
)

func HandleUpdateFile(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed")
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
				c.Redirect(http.StatusBadRequest, url+"?status=failed?reason=not_pdf")
				return
			}
		}

		student, err := store.GetStudent(controlNumber)
		if err != nil {
			fmt.Println("err getting student: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed")
			return
		}

		location, err := utils.FileSaver(c, file, student.Lastname, controlNumber, key)
		if err != nil {
			fmt.Println("err saving file: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed")
			return
		}

		err = store.UpdateFile(key, controlNumber, location)
		if err != nil {
			fmt.Println("err updating file: ", err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusBadRequest, url+"?status=failed")
			return
		}
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
