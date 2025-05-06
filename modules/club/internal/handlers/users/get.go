package users

import (
	"net/http"

	"app/modules/club/internal/handlers/users/results"
	"app/modules/club/utils"

	"github.com/labstack/echo/v4"
)

func (h Handler) Get(c echo.Context) error {
	users, err := h.UserService.Get(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, results.ErrInternalServerError)
	}

	if utils.IsEmpty(users) {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, users)
}
