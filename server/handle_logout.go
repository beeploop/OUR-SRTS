package server

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/config"
)

func HandleLogout(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	_ = strings.Split(referer, "?")[0]

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	cookie, _ := c.Cookie("user")
	c.SetCookie("user", cookie, -1, "/", config.Env.LocalAddr, false, true)

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/")
}
