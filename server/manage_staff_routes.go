package server

import "github.com/gin-gonic/gin"

func HandleManageStaffRoutes(staff *gin.RouterGroup) {

	staff.GET("/", HandleManageStaff)

	staff.POST("/add-staff", HandlePostAddStaff)

    staff.POST("/enable", HandleEnableStaff)

    staff.POST("/disable", HandleDisableStaff)

    staff.POST("/delete", HandleDeleteStaff)

}
