package benefits

import (
	"app/modules/club/domain/interfaces/services"
)

type Handler struct {
	benefitService services.IBenefitServices
}

func NewHandler(benefitService services.IBenefitServices) Handler {
	return Handler{
		benefitService: benefitService,
	}
}
