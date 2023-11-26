package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHome(c *gin.Context) {
	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/auth/login")
}
