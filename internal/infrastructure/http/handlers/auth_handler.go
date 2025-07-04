package handlers

import (
	"log/slog"
	"net/http"

	uc "github.com/beeploop/our-srts/internal/application/usecases/auth"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/web/views/pages/auth"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authUseCase    *uc.UseCase
	sessionManager *session.SessionManager
}

func NewAuthHandler(authUseCase *uc.UseCase, sm *session.SessionManager) *authHandler {
	return &authHandler{
		authUseCase:    authUseCase,
		sessionManager: sm,
	}
}

func (h *authHandler) RenderLogin(c echo.Context) error {
	ctx := c.Request().Context()
	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if !ok {
		toast = ""
	}
	foo := toast.(string)

	page := auth.LoginPage(foo)
	return page.Render(ctx, c.Response().Writer)
}

func (h *authHandler) HandleLogin(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.FormValue("username")
	password := c.FormValue("password")

	admin, err := h.authUseCase.Login(ctx, username, password)
	if err != nil {
		slog.Error("Login Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	adminSession := session.FromDomain(admin)
	if err := h.sessionManager.SetSession(c.Response().Writer, c.Request(), adminSession); err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	return c.Redirect(http.StatusSeeOther, "/app/search")
}

func (h *authHandler) HandleLogout(c echo.Context) error {
	if err := h.sessionManager.ClearSession(c.Response().Writer, c.Request()); err != nil {
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, c.Request().Referer())
	}

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
