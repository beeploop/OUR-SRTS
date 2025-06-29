package handlers

import (
	"net/http"
	"slices"

	"github.com/beeploop/our-srts/internal/application/usecases/reset"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/internal/pkg/utils"
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

	requests, err := h.resetUseCase.GetRequestList(ctx)
	if err != nil {
		page := app.RequestsPage(admin, make([]viewmodel.PasswordResetRequest, 0))
		return page.Render(ctx, c.Response().Writer)
	}

	requestModels := slices.AppendSeq(
		make([]viewmodel.PasswordResetRequest, 0),
		utils.Map(requests, func(request *entities.PasswordResetRequest) viewmodel.PasswordResetRequest {
			return viewmodel.PasswordResetRequestFromDomain(request)
		}),
	)

	page := app.RequestsPage(admin, requestModels)
	return page.Render(ctx, c.Response().Writer)
}

func (h *resetHandler) HandleFulfillRequest(c echo.Context) error {
	ctx := c.Request().Context()

	requestID := c.FormValue("requestID")
	newPassword := c.FormValue("newPassword")
	password := c.FormValue("password")

	if err := h.resetUseCase.FulfillRequest(ctx, requestID, newPassword, password); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/requests")
	}

	return c.Redirect(http.StatusSeeOther, "/app/requests")
}

func (h *resetHandler) HandleRejectRequest(c echo.Context) error {
	ctx := c.Request().Context()

	requestID := c.FormValue("requestID")
	password := c.FormValue("password")

	if err := h.resetUseCase.RejectRequest(ctx, requestID, password); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/requests")
	}

	return c.Redirect(http.StatusSeeOther, "/app/requests")
}
