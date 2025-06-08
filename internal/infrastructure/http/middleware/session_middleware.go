package middleware

import (
	"net/http"

	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

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
