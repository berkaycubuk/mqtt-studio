package utils

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ParseJSON(c echo.Context, payload any) error {
	if c.Request().Body == nil {
		return fmt.Errorf("missing request body")
	}
	
	return json.NewDecoder(c.Request().Body).Decode(payload)
}

func WriteError(c echo.Context, status int, err error) {
	c.JSON(status, map[string]string{"error": err.Error()})
}
