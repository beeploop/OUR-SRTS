package http

import (
	"github.com/beeploop/our-srts/internal/application/usecases/admin"
	"github.com/beeploop/our-srts/internal/application/usecases/program"
	"github.com/beeploop/our-srts/internal/application/usecases/student"
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

func (r *Router) appRouterHandler(g *echo.Group) {
	sessionManager := session.NewSessionManager([]byte("secret"))

	g.Use(middleware.EnsureLoggedIn(sessionManager))
	g.Use(middleware.SessionMiddleware(sessionManager))

	adminRepo := repositories.NewAdminRepository(r.db)
	adminUseCase := admin.NewUseCase(adminRepo)

	studentRepo := repositories.NewStudentRepository(r.db)
	studentUseCase := student.NewUseCase(studentRepo)

	programRepo := repositories.NewProgramRepository(r.db)
	programUseCase := program.NewUseCase(programRepo)

	studentHandler := handlers.NewStudentHandler(studentUseCase, programUseCase)
	accountHandler := handlers.NewAccountHandler(adminUseCase)
	resetHandler := handlers.NewResetHandler()

	g.GET("/search", studentHandler.RenderSearch)
	g.GET("/add-student", studentHandler.RenderAddStudentPage)
	g.POST("/add-student", studentHandler.HandleAddStudent)
	g.GET("/manage-staff", accountHandler.RenderManageStaffPage)
	g.POST("/manage-staff", accountHandler.HandleAddAccount)
	g.POST("/manage-staff/:id/delete", accountHandler.HandleDeleteAccount)
	g.POST("/manage-staff/:id/disable", accountHandler.HandleDisableAccount)
	g.GET("/requests", resetHandler.RenderRequestsPage)
}
