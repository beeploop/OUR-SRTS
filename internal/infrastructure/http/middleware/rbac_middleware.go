package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/labstack/echo/v4"
)

func RBACMiddleware(sm *session.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			admin, ok := sm.GetAdmin(c.Request())
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}

			model := viewmodel.AdminModelFromSession(admin)
			_ = model

			url := c.Request().URL.String()
			trimmedURL := strings.TrimPrefix(url, "/app")
			trimmedURL = strings.Split(trimmedURL, "?")[0]
			protectedRoutes := []string{"/manage-staff", "/requests"}

			if model.Role == string(entities.ROLE_STAFF) {
				if slices.Contains(protectedRoutes, trimmedURL) {
					return c.Redirect(http.StatusSeeOther, "/app/search")
				}
			}

			return next(c)
		}
	}
}
