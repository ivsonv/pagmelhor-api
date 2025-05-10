package partners

import (
	requests "app/modules/club/domain/dto/requests/partners"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Create(c echo.Context) error {
	request := requests.CreatePartnerRequestDto{}

	errors, err := utils.Bind(c, &request)
	if err != nil {
		log.Printf("Bind error in handler create partner: %v", err)
		return c.JSON(http.StatusBadRequest, results.NewErrorWithDetails("BIND_CREATE_PARTNER_ERROR", "Erro ao processar a requisição", errors))
	}

	ctx, cancel := utils.GetContext(c.Request().Context())
	defer cancel()

	result := h.partnerService.Create(ctx, request)
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusCreated, result.Value)
}
