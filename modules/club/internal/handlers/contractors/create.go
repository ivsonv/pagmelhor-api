package contractors

import (
	"net/http"

	requests "app/modules/club/domain/dto/requests/contractors"

	"github.com/labstack/echo/v4"
)

func (h Handler) Create(c echo.Context) error {
	var request requests.ContractorRequestDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result := h.contractorService.Create(c.Request().Context(), request)
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusCreated, result.Value)
}
