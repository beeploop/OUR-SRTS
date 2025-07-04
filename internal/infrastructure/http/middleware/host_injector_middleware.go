package middleware

import (
	"context"

	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/labstack/echo/v4"
)

func HostInjector(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		host := c.Request().Host
		ctx := context.WithValue(c.Request().Context(), contextkeys.HostKey, host)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
