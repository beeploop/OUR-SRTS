package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/registrar/config"
)

func HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	cookie, _ := c.Cookie("user")
	c.SetCookie("user", cookie, -1, "/", config.Env.Ip, false, true)

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/")
}
