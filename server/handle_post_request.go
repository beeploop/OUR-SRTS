package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/registrar/store"
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
		c.Redirect(http.StatusSeeOther, url+"?status=failed")
		return
	}

	err = store.NewRequest(request.Username)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed")
		return
	}

	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/auth/login?status=request_sent")
}
