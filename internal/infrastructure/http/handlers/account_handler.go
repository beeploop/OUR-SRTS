package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"slices"

	"github.com/beeploop/our-srts/internal/application/usecases/admin"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	adminUseCase   *admin.UseCase
	sessionManager *session.SessionManager
}

func NewAccountHandler(
	adminUseCase *admin.UseCase,
	sessionManager *session.SessionManager,
) *accountHandler {
	return &accountHandler{
		adminUseCase:   adminUseCase,
		sessionManager: sessionManager,
	}
}

func (h *accountHandler) RenderManageStaffPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		slog.Error("Session Context", "error", "could not convert session from context to viewmodel.Admin")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	accountModels, err := h.adminUseCase.GetAccounts(ctx)
	if err != nil {
		slog.Error("Get Account Failed", "error", err.Error())
		page := app.ManageStaffPage(admin, make([]viewmodel.Admin, 0))
		return page.Render(ctx, c.Response().Writer)
	}

	accounts := slices.AppendSeq(
		make([]viewmodel.Admin, 0),
		utils.Map(accountModels, func(account *entities.Admin) viewmodel.Admin {
			return viewmodel.AdminFromDomain(account)
		}),
	)

	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if ok {
		ctx = context.WithValue(ctx, contextkeys.ToastKey, toast)
	}

	page := app.ManageStaffPage(admin, accounts)
	return page.Render(ctx, c.Response().Writer)
}

func (h *accountHandler) HandleAddAccount(c echo.Context) error {
	ctx := c.Request().Context()

	fullname := c.FormValue("fullname")
	username := c.FormValue("username")
	password := c.FormValue("password")

	admin := entities.NewAdmin(fullname, username, password, entities.ROLE_STAFF)

	if err := h.adminUseCase.CreateAccount(ctx, admin); err != nil {
		slog.Error("Create Account Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("new staff account created")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		slog.Error("Flash Message", "error", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleDeleteAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.DeleteAccount(ctx, accountID, password); err != nil {
		slog.Error("Delete Account Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("account deleted permanently")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		slog.Error("Flash Message", "error", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleDisableAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.DisableAccount(ctx, accountID, password); err != nil {
		slog.Error("Disable Account Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("account disabled")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		slog.Error("Flash Message", "error", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleEnableAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.EnableAccount(ctx, accountID, password); err != nil {
		slog.Error("Enable Account Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("account re-enabled")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}
