package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleAdminSearchPage(c *gin.Context) {
	html := utils.HtmlParser(
		"admin/search.tmpl",
		"components/header.tmpl",
		"components/sidebar.tmpl",
		"components/searchbar.tmpl",
	)

	user := utils.GetUserInSession(c)
	requestCount := store.CountActiveRequests()

	programs, err := store.GetPrograms()
	if err != nil {
		html.Execute(c.Writer, gin.H{
			"user":         user,
			"students":     []string{},
			"programs":     []string{},
			"requestCount": requestCount,
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":         user,
		"students":     []string{},
		"programs":     programs,
		"requestCount": requestCount,
	})
}
