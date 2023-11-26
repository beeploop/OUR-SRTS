package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/types"
)

func SessionChecker(c *gin.Context) {
	session := sessions.Default(c)
	userSession := session.Get("user")
	_, ok := userSession.(types.User)
	if !ok {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?reason=session_expired")
		return
	}

	c.Next()
}
