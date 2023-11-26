package server

import (
	"github.com/gin-gonic/gin"
	"github.com.BeepLoop/registrar-digitized/utils"
)

func HandleManageProgram(c *gin.Context) {
	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"admin/manage-program.html",
		"components/head.html",
		"components/header.html",
		"components/sidebar.html",
	)

	html.Execute(c.Writer, gin.H{
		"user": user,
	})
}
