package utils

import (
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserInSession(c *gin.Context) *types.User {
	session := sessions.Default(c)
	userSession := session.Get("user")
	user, _ := userSession.(types.User)

	return &user
}
