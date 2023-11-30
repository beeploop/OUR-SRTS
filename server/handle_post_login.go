package server

import (
	"net/http"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func HandlePostLogin(c *gin.Context) {
    logrus.Info("Hit Post login route")
	var input types.Credentials

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		logrus.Info("login error: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed&reason=invalid_form")
		return
	}

	res, err := store.GetCredentials(input.Username)
	if err != nil {
        logrus.Info("logged in with invalid credentials")
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed&reason=wrong_credentials")
		return
	}

	if res.Status == "disabled" {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed&reason=account_disabled")
		return
	}

	err = utils.ValidateCredentials(input.Password, res.Password)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed&reason=wrong_credentials")
		return
	}

	session := sessions.Default(c)
	session.Set("user", types.User{
		Fullname: res.Fullname,
		Username: res.Username,
		Role:     res.Role,
	})
	session.Save()

	c.Request.Method = "GET"
	if res.Role == "admin" {
		c.Redirect(http.StatusSeeOther, "/admin/search")
	} else if res.Role == "staff" {
		c.Redirect(http.StatusSeeOther, "/staff/search")
	}
}
