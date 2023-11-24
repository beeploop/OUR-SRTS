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
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	}

	res, err := store.GetCredentials(input.Username)
	if err != nil {
		fmt.Println("can't get credentials: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed")
	}

	if res.Status == "disabled" {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed")
		return
	}

	if input.Username != res.Username || input.Password != res.Password {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=failed")
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
	c.Redirect(http.StatusSeeOther, "/admin/search")
}
