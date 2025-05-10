package benefits

import (
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
	"net/http"
)

type BenefitService struct {
	contractorRepository repository.IContractorRepository
	benefitRepository    repository.IBenefitRepository
	partnerRepository    repository.IPartnerRepository
}

func NewBenefitService(
	contractorRepository repository.IContractorRepository,
	benefitRepository repository.IBenefitRepository,
	partnerRepository repository.IPartnerRepository,
) services.IBenefitServices {
	return &BenefitService{
		contractorRepository: contractorRepository,
		benefitRepository:    benefitRepository,
		partnerRepository:    partnerRepository,
	}
}

var (
	ErrCreateBenefit      = results.NewError("CREATE_BENEFIT_ERROR", "error creating benefit", http.StatusUnprocessableEntity)
	ErrInvalidEntity      = results.NewError("INVALID_BENEFIT_ENTITY", "invalid entity mapping", http.StatusBadRequest)
	ErrTimeoutOrCanceled  = results.NewError("TIMEOUT_OR_CANCELED", "timeout or canceled", http.StatusRequestTimeout)
	ErrPartnerNotFound    = results.NewError("PARTNER_NOT_FOUND", "partner ID not found", http.StatusNotFound)
	ErrContractorNotFound = results.NewError("CONTRACTOR_NOT_FOUND", "contractor ID not found", http.StatusNotFound)
	ErrInternalServer     = results.NewError("INTERNAL_SERVER_ERROR", "internal server error", http.StatusInternalServerError)
)
