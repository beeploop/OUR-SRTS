package handlers

import (
	"net/http"

	"github.com/beeploop/our-srts/internal/application/usecases/reset"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/beeploop/our-srts/web/views/pages/auth"
	"github.com/labstack/echo/v4"
)

type resetHandler struct {
	resetUseCase *reset.UseCase
}

func NewResetHandler(
	resetUseCase *reset.UseCase,
) *resetHandler {
	return &resetHandler{
		resetUseCase: resetUseCase,
	}
}

func (h *resetHandler) RenderRequestResetPage(c echo.Context) error {
	ctx := c.Request().Context()

	page := auth.ResetRequestPage()
	return page.Render(ctx, c.Response().Writer)
}

func (h *resetHandler) HandleRequestReset(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.FormValue("username")

	if err := h.resetUseCase.RequestPasswordReset(ctx, username); err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h *resetHandler) RenderRequestsListPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	page := app.RequestsPage(admin)
	return page.Render(ctx, c.Response().Writer)
}
