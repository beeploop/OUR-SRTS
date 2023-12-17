package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandlePostRequest(c *gin.Context) {

	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	type Request struct {
		Username string `form:"username" bindig:"required"`
	}

	var request Request
	err := c.ShouldBindWith(&request, binding.Form)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=invalid_form")
		return
	}

	err = store.NewRequest(request.Username)
	if err != nil {
		c.Request.Method = "GET"
		if strings.Contains(err.Error(), "Admin") {
			c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=Admin_cannot_submit_a_request")
			return
		}
		c.Redirect(http.StatusSeeOther, url+"?status=failed&reason=user_have_active_request")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/auth/login?status=success")
}
