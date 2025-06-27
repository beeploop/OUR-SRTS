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

	handler := handlers.NewAppHandler(sessionManager, adminUseCase, studentUseCase, programUseCase)

	g.GET("/search", handler.RenderSearch)
	g.GET("/add-student", handler.RenderAddStudentPage)
	g.POST("/add-student", handler.HandleAddStudent)
	g.GET("/manage-staff", handler.RenderManageStaffPage)
	g.GET("/requests", handler.RenderRequestsPage)
}
