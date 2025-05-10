package benefit_items

import (
	requests "app/modules/club/domain/dto/requests/benefit_items"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Create(c echo.Context) error {
	request := requests.CreateBenefitItemRequestDto{}

	errors, err := utils.Bind(c, &request)
	if err != nil {
		log.Printf("Bind error in handlers.benefititems.create: %v", err)
		return c.JSON(http.StatusBadRequest, results.NewErrorWithDetails("BIND_CREATE_BENEFIT_ITEM_ERROR", "Erro ao processar a requisição", errors))
	}

	ctx, cancel := utils.GetContext(c.Request().Context())
	defer cancel()

	result := h.benefitItemService.Create(ctx, request)
	if !result.IsSuccess {
		return c.JSON(result.Error.StatusCode, result.Error)
	}

	return c.JSON(http.StatusCreated, result.Value)
}
