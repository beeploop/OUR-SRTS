package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/registrar/store"
	"github.com/registrar/types"
)

func HandlePostAddStaff(c *gin.Context) {

	var input types.StaffInfo
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		log.Println("err binding: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=failed")
		return
	}

	err = store.AddUser(input)
	if err != nil {
		log.Println("err adding user: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=failed")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=success")

}
