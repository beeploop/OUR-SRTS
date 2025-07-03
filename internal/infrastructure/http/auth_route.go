package http

import (
	"github.com/beeploop/our-srts/internal/application/usecases/auth"
	"github.com/beeploop/our-srts/internal/application/usecases/reset"
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

func (r *Router) authRouteHandler(g *echo.Group) {
	sessionManager := session.NewSessionManager([]byte(r.cfg.SECRET_KEY))

	adminRepo := repositories.NewAdminRepository(r.db)
	authUseCase := auth.NewUseCase(adminRepo)

	resetRepo := repositories.NewPasswordResetRepository(r.db)
	resetUseCase := reset.NewUseCase(adminRepo, resetRepo)

	authHandler := handlers.NewAuthHandler(authUseCase, sessionManager)
	resetHandler := handlers.NewResetHandler(resetUseCase, sessionManager)

	g.GET("/login", authHandler.RenderLogin, middleware.PreventLogin(sessionManager))
	g.POST("/login", authHandler.HandleLogin, middleware.PreventLogin(sessionManager))
	g.POST("/logout", authHandler.HandleLogout)
	g.GET("/reset/request", resetHandler.RenderRequestResetPage)
	g.POST("/reset/request", resetHandler.HandleRequestReset)
}
