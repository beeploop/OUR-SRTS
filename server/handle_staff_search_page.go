package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleStaffSearchPage(c *gin.Context) {
	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"staff/search.tmpl",
		"components/header.tmpl",
		"components/searchbar.tmpl",
	)

	programs, _ := store.GetPrograms()

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"students": []string{},
		"programs": programs,
	})
}
