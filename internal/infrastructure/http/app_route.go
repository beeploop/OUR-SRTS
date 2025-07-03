package http

import (
	"github.com/beeploop/our-srts/internal/application/usecases/admin"
	"github.com/beeploop/our-srts/internal/application/usecases/program"
	"github.com/beeploop/our-srts/internal/application/usecases/reset"
	"github.com/beeploop/our-srts/internal/application/usecases/student"
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

func (r *Router) appRouterHandler(g *echo.Group) {
	sessionManager := session.NewSessionManager([]byte(r.cfg.SECRET_KEY))

	g.Use(middleware.EnsureLoggedIn(sessionManager))
	g.Use(middleware.SessionMiddleware(sessionManager))
	g.Use(middleware.RBACMiddleware(sessionManager))
	g.Use(middleware.HostInjector)

	documentRepo := repositories.NewDocumentRepository(r.db)
	documentTypeRepo := repositories.NewDocumentTypeRepository(r.db)

	adminRepo := repositories.NewAdminRepository(r.db)
	adminUseCase := admin.NewUseCase(adminRepo)

	studentRepo := repositories.NewStudentRepository(r.db)
	studentUseCase := student.NewUseCase(studentRepo, documentRepo, documentTypeRepo, r.storage)

	programRepo := repositories.NewProgramRepository(r.db)
	programUseCase := program.NewUseCase(programRepo)

	resetRepo := repositories.NewPasswordResetRepository(r.db)
	resetUseCase := reset.NewUseCase(adminRepo, resetRepo)

	studentHandler := handlers.NewStudentHandler(studentUseCase, programUseCase)
	accountHandler := handlers.NewAccountHandler(adminUseCase)
	resetHandler := handlers.NewResetHandler(resetUseCase)

	g.GET("/search", studentHandler.RenderSearch)
	g.GET("/search/:controlNumber", studentHandler.RenderStudentPage)
	g.POST("/search/:controlNumber/update", studentHandler.HandleUpdateStudent)
	g.POST("/search/:controlNumber/upload", studentHandler.HandleUploadDocument)
	g.POST("/search/:controlNumber/reupload", studentHandler.HandleReuploadDocument)
	g.GET("/add-student", studentHandler.RenderAddStudentPage)
	g.POST("/add-student", studentHandler.HandleAddStudent)
	g.GET("/manage-staff", accountHandler.RenderManageStaffPage)
	g.POST("/manage-staff", accountHandler.HandleAddAccount)
	g.POST("/manage-staff/:id/delete", accountHandler.HandleDeleteAccount)
	g.POST("/manage-staff/:id/disable", accountHandler.HandleDisableAccount)
	g.POST("/manage-staff/:id/enable", accountHandler.HandleEnableAccount)
	g.GET("/requests", resetHandler.RenderRequestsListPage)
	g.POST("/requests/:id/fulfill", resetHandler.HandleFulfillRequest)
	g.POST("/requests/:id/reject", resetHandler.HandleRejectRequest)
}
