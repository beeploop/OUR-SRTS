package handlers

import (
	"context"

	"github.com/beeploop/our-srts/web/views/pages/index"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) RenderIndex(c echo.Context) error {
	page := index.IndexPage()
	return page.Render(context.Background(), c.Response().Writer)
}
