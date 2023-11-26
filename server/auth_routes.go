package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAuthRoutes(auth *gin.RouterGroup) {

	auth.GET("/", func(c *gin.Context) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login")
	})

	auth.GET("/login", HandleGetLogin)

	auth.POST("/login", HandlePostLogin)

	auth.POST("/logout", HandleLogout)
}
