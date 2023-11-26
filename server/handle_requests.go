package server

import (
	"github.com/gin-gonic/gin"
	"github.com/registrar/store"
	"github.com/registrar/utils"
)

func HandleRequests(c *gin.Context) {
	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"admin/requests.html",
		"components/header.html",
		"components/sidebar.html",
	)

	requests, err := store.GetRequests()
	if err != nil {
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"requests": requests,
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"requests": requests,
	})
}
