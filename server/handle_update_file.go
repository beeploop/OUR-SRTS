package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/registrar/store"
	"github.com/registrar/utils"
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

	for key := range form.File {
		file := form.File[key][0]

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
