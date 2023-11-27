package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleRequestFulfill(c *gin.Context) {

	user := utils.GetUserInSession(c)
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	type Accept struct {
		RequestId   string `form:"requestId" binding:"required"`
		Password    string `form:"password" binding:"required"`
		NewPassword string `form:"newPassword" binding:"required"`
	}

	var accept Accept
	err := c.ShouldBindWith(&accept, binding.Form)
	if err != nil {
		fmt.Println("err binding: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_request")
		return
	}

	credential, err := store.GetCredentials(user.Username)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=cant_get_credentials")
		return
	}

	fmt.Println("credential: ", credential)
	fmt.Println("accept: ", accept)

	if credential.Password != accept.Password {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_password")
		return
	}

	err = store.FulfillRequest(accept.RequestId, accept.NewPassword)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=database_error")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, url+"?status=success")
}
