package utils

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/types"
)

func GetUserInSession(c *gin.Context) *types.User {
	session := sessions.Default(c)
	userSession := session.Get("user")
	user, _ := userSession.(types.User)

	fmt.Println("user: ", user)
	return &user
}
