package healthz

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Get(c echo.Context) error {
	result := h.HealthzService.Get(c.Request().Context())
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusOK, result.Value)
}
