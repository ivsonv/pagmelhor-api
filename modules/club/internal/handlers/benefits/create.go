package benefits

import (
	"app/modules/club/domain/dto/requests/benefits"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Create(c echo.Context) error {
	request := benefits.CreateBenefitRequestDto{}

	errors, err := utils.Bind(c, &request)
	if err != nil {
		log.Printf("Bind error in handler create benefit: %v", err)
		return c.JSON(http.StatusBadRequest, results.NewErrorWithDetails("BIND_CREATE_BENEFIT_ERROR", "Erro ao processar a requisição", errors))
	}

	ctx, cancel := utils.GetContext(c.Request().Context())
	defer cancel()

	result := h.benefitService.Create(ctx, request)
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusCreated, result.Value)
}
