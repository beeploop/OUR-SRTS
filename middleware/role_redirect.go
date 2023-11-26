package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/registrar/types"
)

func RoleChecker(c *gin.Context) {

	session := sessions.Default(c)
	userSession := session.Get("user")
	user, ok := userSession.(types.User)
	if !ok {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?reason=session_expired")
		return
	}

	url := strings.Split(c.Request.URL.String(), "/")[1]

	if user.Role == "admin" && url != "admin" {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/search")
		return
	}

	if user.Role == "staff" && url != "staff" {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/staff/search")
		return
	}

	c.Next()
}
