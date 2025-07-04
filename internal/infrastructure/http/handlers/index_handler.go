package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type indexHandler struct{}

func NewIndexHandler() *indexHandler {
	return &indexHandler{}
}

func (h *indexHandler) RenderIndex(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/app/search")
}
