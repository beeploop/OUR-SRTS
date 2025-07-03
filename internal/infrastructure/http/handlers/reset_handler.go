package handlers

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeploop/our-srts/internal/application/usecases/reset"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/beeploop/our-srts/web/views/pages/auth"
	"github.com/labstack/echo/v4"
)

type resetHandler struct {
	resetUseCase   *reset.UseCase
	sessionManager *session.SessionManager
}

func NewResetHandler(
	resetUseCase *reset.UseCase,
	sessionManager *session.SessionManager,
) *resetHandler {
	return &resetHandler{
		resetUseCase:   resetUseCase,
		sessionManager: sessionManager,
	}
}

func (h *resetHandler) RenderRequestResetPage(c echo.Context) error {
	ctx := c.Request().Context()

	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if !ok {
		toast = ""
	}
	err := toast.(string)

	page := auth.ResetRequestPage(err)
	return page.Render(ctx, c.Response().Writer)
}

func (h *resetHandler) HandleRequestReset(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.FormValue("username")

	if err := h.resetUseCase.RequestPasswordReset(ctx, username); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, c.Request().Referer())
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

	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if ok {
		ctx = context.WithValue(ctx, contextkeys.ToastKey, toast)
	}

	page := app.RequestsPage(admin, requestModels)
	return page.Render(ctx, c.Response().Writer)
}

func (h *resetHandler) HandleFulfillRequest(c echo.Context) error {
	ctx := c.Request().Context()

	requestID := c.FormValue("requestID")
	newPassword := c.FormValue("newPassword")
	password := c.FormValue("password")

	if err := h.resetUseCase.FulfillRequest(ctx, requestID, newPassword, password); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/requests")
	}

	toast := viewmodel.NewSuccessToast("account password reset successful")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/requests")
}

func (h *resetHandler) HandleRejectRequest(c echo.Context) error {
	ctx := c.Request().Context()

	requestID := c.FormValue("requestID")
	password := c.FormValue("password")

	if err := h.resetUseCase.RejectRequest(ctx, requestID, password); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/requests")
	}

	toast := viewmodel.NewSuccessToast("password reset rejected")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/requests")
}
