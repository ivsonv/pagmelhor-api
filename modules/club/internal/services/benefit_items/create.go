package benefit_items

import (
	requests "app/modules/club/domain/dto/requests/benefit_items"
	responses "app/modules/club/domain/dto/responses/benefit_items"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
	"context"
	"log"
)

func (s *BenefitItemService) Create(ctx context.Context, request requests.CreateBenefitItemRequestDto) results.Result[responses.CreateBenefitItemResponseDto] {
	benefitItem := request.ToMapEntity()
	if benefitItem == nil {
		log.Printf("failed to mapping entity services.benefitItem.create: %v", ErrInvalidEntity)
		return results.Failure[responses.CreateBenefitItemResponseDto](ErrInvalidEntity)
	}

	if err := s.benefitItemRepository.Create(ctx, benefitItem); err != nil {
		if utils.IsTimeout(ctx) {
			return results.Failure[responses.CreateBenefitItemResponseDto](ErrTimeoutOrCanceled)
		}
		return results.Failure[responses.CreateBenefitItemResponseDto](ErrCreateBenefitItem)
	}

	return results.Success(responses.CreateBenefitItemResponseDto{
		ID: benefitItem.ID,
	})
}
