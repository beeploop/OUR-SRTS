package handlers

import (
	"net/http"

	"github.com/beeploop/our-srts/internal/application/usecases"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/web/views/pages/auth"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authUseCase    *usecases.AuthUseCase
	sessionManager *session.SessionManager
}

func NewAuthHandler(authUseCase *usecases.AuthUseCase, sm *session.SessionManager) *authHandler {
	return &authHandler{
		authUseCase:    authUseCase,
		sessionManager: sm,
	}
}

func (h *authHandler) RenderLogin(c echo.Context) error {
	page := auth.LoginPage()
	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (h *authHandler) HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	admin, err := h.authUseCase.Login(username, password)
	if err != nil {
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
		return c.Redirect(http.StatusSeeOther, c.Request().Referer())
	}

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
