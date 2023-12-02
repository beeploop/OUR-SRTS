package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleRequests(c *gin.Context) {
	user := utils.GetUserInSession(c)
    requestCount := store.CountActiveRequests()

	html := utils.HtmlParser(
		"admin/requests.html",
		"components/header.html",
		"components/sidebar.html",
		"components/fulfill-request-modal.html",
		"components/reject-request-modal.html",
	)

	requests, err := store.GetRequests()
	if err != nil {
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"requests": requests,
            "requestCount": requestCount,
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"requests": requests,
        "requestCount": requestCount,
	})
}
