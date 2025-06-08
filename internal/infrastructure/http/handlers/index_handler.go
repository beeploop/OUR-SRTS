package handlers

import (
	"context"

	"github.com/beeploop/our-srts/web/views/pages/index"
	"github.com/labstack/echo/v4"
)

type indexHandler struct{}

func NewIndexHandler() *indexHandler {
	return &indexHandler{}
}

func (h *indexHandler) RenderIndex(c echo.Context) error {
	page := index.IndexPage()
	return page.Render(context.Background(), c.Response().Writer)
}
