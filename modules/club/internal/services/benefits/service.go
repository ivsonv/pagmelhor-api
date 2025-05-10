package benefits

import (
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
	"net/http"
)

type BenefitService struct {
	benefitRepository repository.IBenefitRepository
}

func NewBenefitService(benefitRepository repository.IBenefitRepository) services.IBenefitServices {
	return &BenefitService{
		benefitRepository: benefitRepository,
	}
}

var (
	ErrCreateBenefit     = results.NewError("CREATE_BENEFIT_ERROR", "error creating benefit", http.StatusInternalServerError)
	ErrInvalidEntity     = results.NewError("INVALID_BENEFIT_ENTITY", "invalid entity mapping", http.StatusBadRequest)
	ErrTimeoutOrCanceled = results.NewError("TIMEOUT_OR_CANCELED", "timeout or canceled", http.StatusRequestTimeout)
)
