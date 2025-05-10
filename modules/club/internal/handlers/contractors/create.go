package contractors

import (
	"log"
	"net/http"

	requests "app/modules/club/domain/dto/requests/contractors"
	"app/modules/club/domain/results"
	"app/modules/club/utils"

	"github.com/labstack/echo/v4"
)

func (h Handler) Create(c echo.Context) error {
	request := requests.CreateContractorRequestDto{}

	errors, err := utils.Bind(c, &request)
	if err != nil {
		log.Printf("Bind error in handlers.contractors.create: %v", err)
		return c.JSON(http.StatusBadRequest, results.NewErrorWithDetails("BIND_CREATE_CONTRACTOR_ERROR", "Erro ao processar a requisição", errors))
	}

	ctx, cancel := utils.GetContext(c.Request().Context())
	defer cancel()

	result := h.contractorService.Create(ctx, request)
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusCreated, result.Value)
}
