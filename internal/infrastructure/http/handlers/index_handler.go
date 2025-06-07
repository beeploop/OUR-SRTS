package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) RenderIndex(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
