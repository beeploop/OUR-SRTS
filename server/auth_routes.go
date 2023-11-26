package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/registrar/utils"
)

func HandleAuthRoutes(auth *gin.RouterGroup) {

	auth.GET("/", func(c *gin.Context) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login")
	})

	auth.GET("/login", HandleGetLogin)

	auth.POST("/login", HandlePostLogin)

	auth.POST("/logout", HandleLogout)

	auth.GET("/request", func(c *gin.Context) {
		html := utils.HtmlParser(
			"request-page.html",
			"components/header.html",
			// "components/request.html",
		)

		html.Execute(c.Writer, gin.H{})
	})
}
