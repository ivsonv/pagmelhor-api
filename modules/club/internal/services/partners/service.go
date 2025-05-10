package partners

import (
	"net/http"

	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
)

type PartnerService struct {
	partnerRepository repository.IPartnerRepository
}

func NewPartnerService(partnerRepository repository.IPartnerRepository) services.IPartnerServices {
	return &PartnerService{partnerRepository: partnerRepository}
}

var (
	ErrCreatePartner     = results.NewError("CREATE_PARTNER_ERROR", "error creating partner", http.StatusUnprocessableEntity)
	ErrInvalidEntity     = results.NewError("INVALID_PARTNER_ENTITY", "invalid to mapentity", http.StatusBadRequest)
	ErrTimeoutOrCanceled = results.NewError("TIMEOUT_OR_CANCELED_PARTNER_ERROR", "operation partner timeout or canceled", http.StatusRequestTimeout)
	ErrCpfCnpjDuplicated = results.NewError("CPF_CNPJ_PARTNER_DUPLICATED", "cpf/cnpj already exists", http.StatusConflict)
	ErrEmailDuplicated   = results.NewError("EMAIL_PARTNER_DUPLICATED", "email already exists", http.StatusConflict)
	ErrSlugDuplicated    = results.NewError("SLUG_PARTNER_DUPLICATED", "slug already exists", http.StatusConflict)
	ErrInternalServer    = results.NewError("INTERNAL_SERVER_ERROR", "internal server error", http.StatusInternalServerError)
)
