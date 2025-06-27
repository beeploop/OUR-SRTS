package handlers

import (
	"net/http"

	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/labstack/echo/v4"
)

type resetHandler struct{}

func NewResetHandler() *resetHandler {
	return &resetHandler{}
}

func (h *resetHandler) RenderRequestsPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	page := app.RequestsPage(admin)
	return page.Render(c.Request().Context(), c.Response().Writer)
}
