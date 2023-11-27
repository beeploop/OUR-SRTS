package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAdminRoutes(admin *gin.RouterGroup) {

	admin.GET("/", func(c *gin.Context) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/search")
	})

	admin.GET("/search", HandleAdminSearchPage)

	admin.GET("/add-student", HandleGetAddStudent)

	admin.POST("/add-student", HandlePostAddStudent)

	admin.POST("/update-student", HandleUpdateStudent)

	admin.POST("/add-staff", HandlePostAddStaff)

	admin.GET("/manage-program", HandleManageProgram)

	admin.GET("/programs", HandleGetPrograms)

	request := admin.Group("/request")
	HandleRequestRoutes(request)

	manageStaff := admin.Group("/manage-staff")
	HandleManageStaffRoutes(manageStaff)

}
