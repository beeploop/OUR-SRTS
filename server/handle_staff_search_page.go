package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleStaffSearchPage(c *gin.Context) {
	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"staff/search.html",
		"components/header.html",
		"components/searchbar.html",
	)

	programs, _ := store.GetPrograms()

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"students": []string{},
		"programs": programs,
	})
}
