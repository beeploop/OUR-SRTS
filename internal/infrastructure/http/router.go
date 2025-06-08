package http

import (
	"encoding/gob"

	"github.com/beeploop/our-srts/internal/config"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg  *config.Config
	db   *sqlx.DB
	Echo *echo.Echo
}

func NewRouter(cfg *config.Config, db *sqlx.DB) *Router {
	r := echo.New()

	gob.Register(session.SessionModel{})

	router := &Router{
		cfg:  cfg,
		db:   db,
		Echo: r,
	}

	router.registerRoutes()

	return router
}

func (r *Router) registerRoutes() {

	r.Echo.Use(middleware.RemoveTrailingSlash())

	r.Echo.Static("/assets", "web/assets/")

	indexRoute := r.Echo.Group("/")
	r.indexRouteHandler(indexRoute)

	authRoute := r.Echo.Group("/auth")
	r.authRouteHandler(authRoute)

	appRoute := r.Echo.Group("/app")
	r.appRouterHandler(appRoute)
}
