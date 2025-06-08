package http

import (
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

	studentRepo := repositories.NewStudentRepository(r.db)
	studentUseCase := student.NewUseCase(studentRepo)

	handler := handlers.NewAppHandler(sessionManager, studentUseCase)

	g.GET("/search", handler.RenderSearch)
	g.GET("/add-student", handler.RenderAddStudentPage)
	g.GET("/manage-staff", handler.RenderManageStaffPage)
	g.GET("/requests", handler.RenderRequestsPage)
}
