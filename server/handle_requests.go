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
		"admin/requests.tmpl",
		"components/header.tmpl",
		"components/sidebar.tmpl",
		"components/fulfill-request-modal.tmpl",
		"components/reject-request-modal.tmpl",
	)

	requests, err := store.GetRequests()
	if err != nil {
		html.Execute(c.Writer, gin.H{
			"user":         user,
			"requests":     requests,
			"requestCount": requestCount,
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":         user,
		"requests":     requests,
		"requestCount": requestCount,
	})
}
