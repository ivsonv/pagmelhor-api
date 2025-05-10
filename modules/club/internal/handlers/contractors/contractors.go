package contractors

import (
	"app/modules/club/domain/interfaces/services"
)

type Handler struct {
	contractorService services.IContractorServices
}

func NewHandler(contractorService services.IContractorServices) Handler {
	return Handler{contractorService: contractorService}
}
