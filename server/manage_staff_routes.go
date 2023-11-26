package server

import "github.com/gin-gonic/gin"

func HandleManageStaffRoutes(staff *gin.RouterGroup) {

	staff.GET("/", HandleManageStaff)

	staff.POST("/add-staff", HandlePostAddStaff)

    staff.POST("/disable", HandleDisableStaff)

}
