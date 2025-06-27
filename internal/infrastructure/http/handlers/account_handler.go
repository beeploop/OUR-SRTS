package handlers

import (
	"net/http"
	"slices"

	"github.com/beeploop/our-srts/internal/application/usecases/admin"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	adminUseCase *admin.UseCase
}

func NewAccountHandler(
	adminUseCase *admin.UseCase,
) *accountHandler {
	return &accountHandler{
		adminUseCase: adminUseCase,
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
		return page.Render(c.Request().Context(), c.Response().Writer)
	}

	accounts := slices.AppendSeq(
		make([]viewmodel.Admin, 0),
		utils.Map(accountModels, func(account *entities.Admin) viewmodel.Admin {
			return viewmodel.AdminFromDomain(account)
		}),
	)

	page := app.ManageStaffPage(admin, accounts)
	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (h *accountHandler) HandleAddAccount(c echo.Context) error {
	ctx := c.Request().Context()

	fullname := c.FormValue("fullname")
	username := c.FormValue("username")
	password := c.FormValue("password")

	admin := entities.NewAdmin(fullname, username, password, entities.ROLE_STAFF)

	if err := h.adminUseCase.CreateAccount(ctx, admin); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleDeleteAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.DeleteAccount(ctx, accountID, password); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleDisableAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.DisableAccount(ctx, accountID, password); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}

func (h *accountHandler) HandleEnableAccount(c echo.Context) error {
	ctx := c.Request().Context()

	accountID := c.FormValue("accountID")
	password := c.FormValue("password")

	if err := h.adminUseCase.EnableAccount(ctx, accountID, password); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
	}

	return c.Redirect(http.StatusSeeOther, "/app/manage-staff")
}
