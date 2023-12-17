package server

import (
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleManageProgram(c *gin.Context) {
	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"admin/manage-program.tmpl",
		"components/head.tmpl",
		"components/header.tmpl",
		"components/sidebar.tmpl",
	)

	html.Execute(c.Writer, gin.H{
		"user": user,
	})
}
