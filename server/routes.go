package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/registrar/middleware"
	"github.com/registrar/store"
)

func RegisterRoutes() {

	Router.GET("/", HandleHome)

	auth := Router.Group("/auth")
	HandleAuthRoutes(auth)

	staff := Router.Group("/staff", middleware.RoleChecker)
	HandleStaffRoutes(staff)

	admin := Router.Group("/admin", middleware.RoleChecker)
	HandleAdminRoutes(admin)

	student := Router.Group("/student", middleware.SessionChecker)
	HandleStudentRoutes(student)

	Router.POST("/upload", HandleUpload)

	Router.POST("/request", func(c *gin.Context) {
		referer := c.Request.Header.Get("Referer")
		url := strings.Split(referer, "?")[0]

		type Request struct {
			Username string `form:"username" bindig:"required"`
		}

		var request Request
		err := c.ShouldBindWith(&request, binding.Form)
		if err != nil {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, url+"?status=failed")
			return
		}

		fmt.Println("request: ", request)
		err = store.NewRequest(request.Username)
		if err != nil {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, url+"?status=failed")
			return
		}

		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/auth/login?status=request_sent")
	})

}
