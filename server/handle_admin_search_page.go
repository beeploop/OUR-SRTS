package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
)

func HandleAdminSearchPage(c *gin.Context) {
	html := utils.HtmlParser(
		"admin/search.html",
		"components/header.html",
		"components/sidebar.html",
		"components/search.html",
	)

	user := utils.GetUserInSession(c)
	session := sessions.Default(c)
	students := session.Get("search-result")

	programs, err := store.GetPrograms()
	if err != nil {
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"students": students,
			"programs": []string{},
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"students": students,
		"programs": programs,
	})
}
