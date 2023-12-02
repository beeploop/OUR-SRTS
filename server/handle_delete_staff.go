package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func HandleDeleteStaff(c *gin.Context) {

	type DeleteStaff struct {
		Staff    string `form:"staff" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	user := utils.GetUserInSession(c)
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	var input DeleteStaff
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
        logrus.Warn("err binding form: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, url+"?status=failed&reason=invalid_input")
		return
	}

	if input.Staff == user.Username {
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, url+"?status=failed&reason=invalid_input")
		return
	}

	credential, err := store.GetCredentials(user.Username)
	if err != nil {
        logrus.Warn("err getting credential: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, url+"?status=failed&reason=invalid_input")
		return
	}

	err = utils.ValidateCredentials(input.Password, credential.Password)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, url+"?status=failed&reason=invalid_password")
		return
	}

	err = store.DeleteStaff(input.Staff)
	if err != nil {
        logrus.Warn("err deleting staff: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, url+"?status=failed&reason=invalid_input")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusMovedPermanently, url+"?status=success")

}
