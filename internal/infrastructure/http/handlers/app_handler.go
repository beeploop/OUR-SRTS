package handlers

import (
	"errors"

	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/labstack/echo/v4"
)

type appHandler struct {
	sm *session.SessionManager
}

func NewAppHandler(sm *session.SessionManager) *appHandler {
	return &appHandler{
		sm: sm,
	}
}

func (h *appHandler) RenderSearch(c echo.Context) error {
	admin, ok := h.sm.GetAdmin(c.Request())
	if !ok {
		return errors.New("invalid session")
	}

	vm := viewmodel.Admin{
		ID:       admin.ID,
		Fullname: admin.Fullname,
		Username: admin.Username,
		Role:     admin.Role,
	}

	page := app.SearchPage(vm)
	return page.Render(c.Request().Context(), c.Response().Writer)
}
