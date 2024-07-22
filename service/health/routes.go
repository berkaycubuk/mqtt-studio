package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *echo.Group) {
	router.GET("/health", h.handleHealth)
}

func (h *Handler) handleHealth(c echo.Context) error {
	return c.String(http.StatusOK, "Hi!")
}
