package benefit_items

import (
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
	"net/http"
)

type BenefitItemService struct {
	benefitItemRepository repository.IBenefitItemRepository
	benefitRepository     repository.IBenefitRepository
}

func NewBenefitItemService(
	benefitItemRepository repository.IBenefitItemRepository,
	benefitRepository repository.IBenefitRepository,
) services.IBenefitItemServices {
	return &BenefitItemService{
		benefitItemRepository: benefitItemRepository,
		benefitRepository:     benefitRepository,
	}
}

var (
	ErrCreateBenefitItem = results.NewError("error creating benefit item", "error creating benefit item", http.StatusUnprocessableEntity)
	ErrInvalidEntity     = results.NewError("invalid entity mapping", "invalid entity mapping", http.StatusBadRequest)
	ErrTimeoutOrCanceled = results.NewError("timeout or canceled", "timeout or canceled", http.StatusRequestTimeout)
	ErrBenefitNotFound   = results.NewError("benefit not found", "benefit not found", http.StatusNotFound)
	ErrInternalServer    = results.NewError("internal server error", "internal server error", http.StatusInternalServerError)
)
