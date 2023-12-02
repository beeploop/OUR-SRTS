package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleAdminPostSearch(c *gin.Context) {

	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"admin/search.html",
		"components/header.html",
		"components/sidebar.html",
		"components/searchbar.html",
	)

	programs, err := store.GetPrograms()
	if err != nil {
		c.Request.Method = "GET"
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"students": []types.Student{},
			"programs": programs,
		})
		return
	}

	var input types.SearchData
	err = c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		c.Request.Method = "GET"
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"students": []types.Student{},
			"programs": programs,
		})
		return
	}

	if input.SearchTerm == "" {
		c.Request.Method = "GET"
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"students": []types.Student{},
			"programs": programs,
		})
		return
	}

	students, err := store.SearchStudent(input)
	if err != nil {
		c.Request.Method = "GET"
		html.Execute(c.Writer, gin.H{
			"user":     user,
			"students": []string{},
			"programs": programs,
		})
		return
	}

	c.Request.Method = "GET"
	html.Execute(c.Writer, gin.H{
		"user":     user,
		"students": students,
		"programs": programs,
	})
}
