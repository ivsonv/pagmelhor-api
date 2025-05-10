package benefit_items

import (
	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
)

type BenefitItemService struct {
	benefitItemRepository repository.IBenefitItemRepository
}

func NewBenefitItemService(benefitItemRepository repository.IBenefitItemRepository) services.IBenefitItemServices {
	return &BenefitItemService{
		benefitItemRepository: benefitItemRepository,
	}
}

var (
	ErrCreateBenefitItem = results.NewError("error creating benefit item", "error creating benefit item", 500)
	ErrInvalidEntity     = results.NewError("invalid entity mapping", "invalid entity mapping", 400)
	ErrTimeoutOrCanceled = results.NewError("timeout or canceled", "timeout or canceled", 408)
)
