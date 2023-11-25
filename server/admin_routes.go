package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/registrar/store"
)

func HandleAdminRoutes(admin *gin.RouterGroup) {

	admin.GET("/", func(c *gin.Context) {
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/search")
	})

	admin.GET("/search", HandleAdminSearchPage)

	admin.GET("/add-student", HandleGetAddStudent)

	admin.POST("/add-student", HandlePostAddStudent)

	admin.GET("/manage-staff", HandleManageStaff)

	admin.POST("/add-staff", HandlePostAddStaff)

	admin.GET("/manage-program", HandleManageProgram)

	admin.GET("/requests", HandleRequests)

	admin.POST("/upload", HandleUpload)

	admin.GET("/programs", func(c *gin.Context) {
		list, err := store.GetProgramsAndMajors()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "data": nil})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": list})
	})

}
