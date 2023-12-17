package server

import (
	"net/http"

	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleAuthRoutes(auth *gin.RouterGroup) {

	auth.GET("/", func(c *gin.Context) {
		logrus.Info("Hit admin route")
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login")
	})

	auth.GET("/login", HandleGetLogin)

	auth.POST("/login", HandlePostLogin)

	auth.POST("/logout", HandleLogout)

	auth.GET("/request", func(c *gin.Context) {
		html := utils.HtmlParser(
			"request-page.tmpl",
			"components/header.tmpl",
		)

		html.Execute(c.Writer, gin.H{})
	})
}
