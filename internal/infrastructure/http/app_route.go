package http

import (
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

func appRouterHandler(g *echo.Group) {
	sessionManager := session.NewSessionManager([]byte("secret"))
	g.Use(middleware.EnsureLoggedIn(sessionManager))

	handler := handlers.NewAppHandler(sessionManager)

	g.GET("/search", handler.RenderSearch)
}
