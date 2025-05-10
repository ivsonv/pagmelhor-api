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
	if err := benefitItem.IsValid(); err != nil {
		log.Printf("ToMapEntity.IsValid.services.benefit_items.create: %v", err.Error())
		return results.Failure[responses.CreateBenefitItemResponseDto](ErrInvalidEntity)
	}

	if err := s.preValidation(ctx, request); err != nil {
		return results.Failure[responses.CreateBenefitItemResponseDto](*err)
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

func (s *BenefitItemService) preValidation(ctx context.Context, request requests.CreateBenefitItemRequestDto) *results.Error {
	exists, err := s.benefitRepository.ExistsById(ctx, request.BenefitID)
	if err != nil {
		return &ErrInternalServer
	}

	if !exists {
		return &ErrBenefitNotFound
	}

	return nil
}
