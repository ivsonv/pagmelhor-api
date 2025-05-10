package benefit_items

import (
	"app/modules/club/domain/interfaces/services"
)

type Handler struct {
	benefitItemService services.IBenefitItemServices
}

func NewHandler(benefitItemService services.IBenefitItemServices) Handler {
	return Handler{
		benefitItemService: benefitItemService,
	}
}
