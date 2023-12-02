package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func HandleUpdateStudent(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
    url := strings.Split(referer, "?")[0]

	var student types.StudentInfo
	err := c.ShouldBindWith(&student, binding.Form)
	if err != nil {
        logrus.Warn("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
		return
	}

	err = store.UpdateStudent(student)
	if err != nil {
        logrus.Warn("err updating student: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=unknown_student")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
