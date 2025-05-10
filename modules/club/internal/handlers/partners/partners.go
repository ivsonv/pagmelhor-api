package partners

import (
	"app/modules/club/domain/interfaces/services"
)

type Handler struct {
	partnerService services.IPartnerServices
}

func NewHandler(partnerService services.IPartnerServices) Handler {
	return Handler{partnerService: partnerService}
}
