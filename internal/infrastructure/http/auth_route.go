package http

import (
	"github.com/beeploop/our-srts/internal/application/usecases"
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func authRouteHandler(g *echo.Group, db *sqlx.DB) {
	sessionManager := session.NewSessionManager([]byte("secret"))

	adminRepo := repositories.NewAdminRepository(db)
	authUseCase := usecases.NewAuthUseCase(adminRepo)

	handler := handlers.NewAuthHandler(authUseCase, sessionManager)

	g.GET("/login", handler.RenderLogin, middleware.PreventLogin(sessionManager))
	g.POST("/login", handler.HandleLogin, middleware.PreventLogin(sessionManager))
	g.POST("/logout", handler.HandleLogout)
}
