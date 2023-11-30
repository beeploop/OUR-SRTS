package server

import (
	"net/http"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func HandlePostAddStaff(c *gin.Context) {

	var input types.StaffInfo
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		logrus.Info("err binding: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=failed&reason=invalid_form")
		return
	}

	if !utils.ValidPassword(input.Password) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=failed&reason=Password_does_not_meet_requirements")
		return
	}

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=failed&reason=server_error")
		return
	}
	input.Password = hash

	err = store.AddUser(input)
	if err != nil {
		logrus.Info("err adding user: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=failed&reason=invalid_staff_info")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusMovedPermanently, "/admin/manage-staff?status=success")

}
