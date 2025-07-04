package http

import (
	"encoding/gob"

	"github.com/beeploop/our-srts/internal/application/interfaces"
	"github.com/beeploop/our-srts/internal/config"
	"github.com/beeploop/our-srts/internal/infrastructure/http/middleware"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

type Router struct {
	cfg     *config.Config
	db      *sqlx.DB
	storage interfaces.Storage
	Echo    *echo.Echo
}

func NewRouter(cfg *config.Config, db *sqlx.DB, storage interfaces.Storage) *Router {
	r := echo.New()

	gob.Register(session.SessionModel{})

	router := &Router{
		cfg:     cfg,
		db:      db,
		storage: storage,
		Echo:    r,
	}

	router.registerRoutes()

	return router
}

func (r *Router) registerRoutes() {
	r.Echo.Use(md.RemoveTrailingSlash())
	r.Echo.Use(middleware.CustomLogger)

	r.Echo.Static("/assets", "web/assets/")
	r.Echo.Static("/uploads", r.cfg.UPLOAD_DIR)

	indexRoute := r.Echo.Group("/")
	r.indexRouteHandler(indexRoute)

	authRoute := r.Echo.Group("/auth")
	r.authRouteHandler(authRoute)

	appRoute := r.Echo.Group("/app")
	r.appRouterHandler(appRoute)
}
