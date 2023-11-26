package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleDisableStaff(c *gin.Context) {
	type DisableStaff struct {
		Staff    string `form:"staff" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	user := utils.GetUserInSession(c)
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	var input DisableStaff
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed?reason=invalid_input")
		return
	}

	if input.Staff == user.Username {
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed?reason=invalid_input")
		return
	}

	credential, err := store.GetCredentials(user.Username)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed?reason=invalid_input")
		return
	}

	if credential.Password != input.Password {
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed?reason=invalid_password")
		return
	}

	err = store.DisableStaff(input.Staff)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusBadRequest, url+"?status=failed?reason=invalid_input")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
