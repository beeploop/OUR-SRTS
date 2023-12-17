package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleGetAddStudent(c *gin.Context) {
	user := utils.GetUserInSession(c)
	requestCount := store.CountActiveRequests()

	programs, err := store.GetProgramsAndMajors()
	if err != nil {
		logrus.Warn("err getting programs: ", err)
	}

	html := utils.HtmlParser(
		"admin/add-student.tmpl",
		"components/header.tmpl",
		"components/sidebar.tmpl",
	)

	html.Execute(c.Writer, gin.H{
		"user":         user,
		"programs":     programs,
		"requestCount": requestCount,
	})
}
