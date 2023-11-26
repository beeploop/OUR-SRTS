package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
)

func HandleStaffSearchPage(c *gin.Context) {
	user := utils.GetUserInSession(c)
	session := sessions.Default(c)
	students := session.Get("search-result")

	html := utils.HtmlParser(
		"staff/search.html",
		"components/header.html",
		"components/search.html",
	)

	programs, _ := store.GetPrograms()

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"students": students,
		"programs": programs,
	})
}
