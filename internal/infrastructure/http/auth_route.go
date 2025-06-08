package http

import (
	"github.com/beeploop/our-srts/internal/application/usecases/auth"
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

func (r *Router) authRouteHandler(g *echo.Group) {
	sessionManager := session.NewSessionManager([]byte("secret"))

	adminRepo := repositories.NewAdminRepository(r.db)
	authUseCase := auth.NewUseCase(adminRepo)

	handler := handlers.NewAuthHandler(authUseCase, sessionManager)

	g.GET("/login", handler.RenderLogin, middleware.PreventLogin(sessionManager))
	g.POST("/login", handler.HandleLogin, middleware.PreventLogin(sessionManager))
	g.POST("/logout", handler.HandleLogout)
}
