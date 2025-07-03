package middleware

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func CustomLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Info(
			"Route Hit",
			slog.String("TIME", time.Now().Format(time.RFC3339)),
			slog.String("IP", c.RealIP()),
			slog.String("URI", c.Request().RequestURI),
			slog.String("METHOD", c.Request().Method),
			slog.String("PROTOCOL", c.Request().Proto),
			slog.String("REFERER", c.Request().Referer()),
			slog.String("USER_AGENT", c.Request().UserAgent()),
		)

		return next(c)
	}
}
