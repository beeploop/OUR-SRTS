package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandlePostAddStudent(c *gin.Context) {
	var info types.StudentInfo
	err := c.ShouldBindWith(&info, binding.Form)
	if err != nil {
		fmt.Println("error binding info: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/add-student?status=failed&reason=invalid_form")
		return
	}

	err = store.AddStudent(info)
	if err != nil {
		fmt.Println("error insert student to db: ", err)
		c.Request.Method = "GET"

		if strings.Contains(err.Error(), "Duplicate entry") {
			c.Redirect(http.StatusSeeOther, "/admin/add-student?status=failed&reason=Control_number_already_exists")
			return
		}
		c.Redirect(http.StatusSeeOther, "/admin/add-student?status=failed&reason=invalid_student_info")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusMovedPermanently, "/admin/add-student?status=success")
}
