package handlers

import (
	"context"
	"fmt"
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
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	accountModels, err := h.adminUseCase.GetAccounts(ctx)
	if err != nil {
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
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("new staff account created")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleDeleteAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.DeleteAccount(ctx, accountID, password); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("account deleted permanently")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleDisableAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.DisableAccount(ctx, accountID, password); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("account disabled")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleEnableAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.EnableAccount(ctx, accountID, password); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			fmt.Println("error setting fash: ", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	toast := viewmodel.NewSuccessToast("account re-enabled")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}
