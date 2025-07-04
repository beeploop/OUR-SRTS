package middleware

import (
	"context"
	"net/http"

	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/labstack/echo/v4"
)

func SessionMiddleware(sm *session.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			admin, ok := sm.GetAdmin(c.Request())
			if !ok {
				c.Redirect(http.StatusSeeOther, "/auth/login")
			}

			model := viewmodel.AdminModelFromSession(admin)

			ctx := context.WithValue(c.Request().Context(), contextkeys.SessionKey, model)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func EnsureLoggedIn(sm *session.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, ok := sm.GetAdmin(c.Request())
			if !ok {
				c.Redirect(http.StatusSeeOther, "/auth/login")
			}

			return next(c)
		}
	}
}

func PreventLogin(sm *session.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, ok := sm.GetAdmin(c.Request())
			if ok {
				c.Redirect(http.StatusSeeOther, "/app/search")
			}

			return next(c)
		}
	}
}
