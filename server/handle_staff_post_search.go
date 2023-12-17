package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func HandleStaffPostSearch(c *gin.Context) {
	user := utils.GetUserInSession(c)

	html := utils.HtmlParser(
		"staff/search.tmpl",
		"components/header.tmpl",
		"components/searchbar.tmpl",
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
		logrus.Warn(err)
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
