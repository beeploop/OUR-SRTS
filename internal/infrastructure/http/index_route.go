package http

import (
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/labstack/echo/v4"
)

func indexRouteHandler(g *echo.Group) {
	handler := handlers.NewIndexHandler()

	g.GET("", handler.RenderIndex)
}
