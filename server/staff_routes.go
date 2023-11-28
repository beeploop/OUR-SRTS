package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleStaffRoutes(staff *gin.RouterGroup) {

	staff.GET("/", func(c *gin.Context) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/staff/search")
	})

	staff.GET("/search", HandleStaffSearchPage)

	staff.POST("/search", HandleStaffPostSearch)
}
