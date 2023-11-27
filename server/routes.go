package server

import (
	"github.com/BeepLoop/registrar-digitized/middleware"
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

	Router.POST("/upload", middleware.SessionChecker, HandleUpload)

	Router.POST("/update", middleware.SessionChecker, HandleUpdateFile)

	Router.POST("/request", HandlePostRequest)
}
