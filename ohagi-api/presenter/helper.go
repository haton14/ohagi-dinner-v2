package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func sendResponse(c echo.Context, code int) error {
	return c.JSON(code, struct {
		Message string `json:"message"`
	}{Message: http.StatusText(code)})
}
