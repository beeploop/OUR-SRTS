package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/registrar/store"
	"github.com/registrar/types"
)

func HandlePostLogin(c *gin.Context) {
	var input types.Credentials

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		log.Println("login error: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed?reason=can't_bind_form")
		return
	}

	res, err := store.GetCredentials(input.Username)
	if err != nil {
		fmt.Println("can't get credentials: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed?reason=wrong_credentials")
	}

	if res.Status == "disabled" {
        fmt.Println("account disabled")
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed?reason=disabled")
		return
	}

	if input.Username != res.Username || input.Password != res.Password {
        fmt.Println("wrong credentials")
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed?reason=wrong_credentials")
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
