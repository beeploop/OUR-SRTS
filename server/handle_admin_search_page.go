package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleAdminSearchPage(c *gin.Context) {
	html := utils.HtmlParser(
		"admin/search.html",
		"components/header.html",
		"components/sidebar.html",
		"components/searchbar.html",
	)

	user := utils.GetUserInSession(c)
    requestCount := store.CountActiveRequests()

	programs, err := store.GetPrograms()
	if err != nil {
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"students": []string{},
			"programs": []string{},
            "requestCount": requestCount,
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":     user,
		"students": []string{},
		"programs": programs,
        "requestCount": requestCount,
	})
}
