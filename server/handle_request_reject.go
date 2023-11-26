package server

import (
	"net/http"
	"strings"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleRequestReject(c *gin.Context) {
	referer := c.Request.Header.Get("Referer")
	url := strings.Split(referer, "?")[0]

	type Reject struct {
		RequestId string `form:"requestId"`
	}

	var reject Reject
	err := c.ShouldBindWith(&reject, binding.Form)
	if err != nil {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, url+"?status=failed?reason=invalid_request")
		return
	}

    err = store.RejectRequest(reject.RequestId)
    if err != nil {
        c.Request.Method = "GET"
        c.Redirect(http.StatusSeeOther, url+"?status=failed?reason=database_error")
        return
    }

    c.Request.Method = "GET"
    c.Redirect(http.StatusSeeOther, url+"?status=success")
}
