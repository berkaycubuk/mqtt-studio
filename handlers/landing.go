package handlers

import (
	"github.com/berkaycubuk/mqtt-studio/views/landing"
	"github.com/labstack/echo/v4"
)

type LandingHandler struct {
}

func (h LandingHandler) Index(c echo.Context) error {
	return render(c, landing.Index())
}
