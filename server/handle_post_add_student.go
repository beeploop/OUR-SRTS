package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/registrar/store"
	"github.com/registrar/types"
)

func HandlePostAddStudent(c *gin.Context) {
	var info types.StudentInfo
	err := c.ShouldBindWith(&info, binding.Form)
	if err != nil {
		fmt.Println("error binding info: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/add-student?status=failed")
		return
	}

	fmt.Println("info: ", info)
	err = store.AddStudent(info)
	if err != nil {
		fmt.Println("error insert student to db: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/add-student?status=failed")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusMovedPermanently, "/admin/add-student?status=success")
}
