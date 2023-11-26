package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
)

func HandleUpdateStudent(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
    url := strings.Split(referer, "?")[0]

	var student types.StudentInfo
	err := c.ShouldBindWith(&student, binding.Form)
	if err != nil {
		fmt.Println("err binding: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed")
		return
	}

	err = store.UpdateStudent(student)
	if err != nil {
        fmt.Println("err update: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
