package users

import (
	"app/modules/club/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Get(c echo.Context) error {
	result := h.UserService.Get(c.Request().Context())
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	if utils.IsEmpty(result.Value) {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, result.Value)
}
