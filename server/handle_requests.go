package server

import (
	"github.com/gin-gonic/gin"
	"github.com/registrar/utils"
)

func HandleRequests(c *gin.Context) {
    user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"admin/requests.html",
		"components/head.html",
		"components/header.html",
		"components/sidebar.html",
	)

	html.Execute(c.Writer, gin.H{
        "user": user,
    })
}
