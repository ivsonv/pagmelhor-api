package benefits

import (
	requests "app/modules/club/domain/dto/requests/benefits"
	responses "app/modules/club/domain/dto/responses/benefits"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
	"context"
	"log"
)

func (s *BenefitService) Create(ctx context.Context, request requests.CreateBenefitRequestDto) results.Result[responses.CreateBenefitResponseDto] {
	benefit := request.ToMapEntity()
	if benefit == nil {
		log.Printf("ToMapEntity services.benefits.create: %v", ErrInvalidEntity)
		return results.Failure[responses.CreateBenefitResponseDto](ErrInvalidEntity)
	}

	if err := s.benefitRepository.Create(ctx, benefit); err != nil {
		if utils.IsTimeout(ctx) {
			return results.Failure[responses.CreateBenefitResponseDto](ErrTimeoutOrCanceled)
		}

		log.Printf("Create services.benefits.create: %v", err)
		return results.Failure[responses.CreateBenefitResponseDto](ErrCreateBenefit)
	}

	return results.Success(responses.CreateBenefitResponseDto{
		ID: benefit.ID,
	})
}
