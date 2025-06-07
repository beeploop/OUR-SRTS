package http

import (
	"github.com/beeploop/our-srts/internal/infrastructure/http/handlers"
	"github.com/labstack/echo/v4"
)

func indexRouteHandler(r *echo.Group) {
	indexHandler := handlers.NewIndexHandler()

	r.GET("", indexHandler.RenderIndex)
}
