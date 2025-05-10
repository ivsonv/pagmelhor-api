package contractors

import (
	"net/http"

	requests "app/modules/club/domain/dto/requests/contractors"
	"app/modules/club/utils"

	"github.com/labstack/echo/v4"
)

func (h Handler) Create(c echo.Context) error {
	request := requests.ContractorRequestDto{}

	detailsErrors, err := utils.Bind(c, &request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, detailsErrors)
	}

	ctx, cancel := utils.GetContext(c.Request().Context())
	defer cancel()

	result := h.contractorService.Create(ctx, request)
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusCreated, result.Value)
}
