package server

import (
	"github.com/registrar/middleware"
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

}
