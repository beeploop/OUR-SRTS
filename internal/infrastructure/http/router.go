package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() *echo.Echo {
	r := echo.New()

	r.Use(middleware.RemoveTrailingSlash())

	indexRoute := r.Group("/")
	indexRouteHandler(indexRoute)

	return r
}
