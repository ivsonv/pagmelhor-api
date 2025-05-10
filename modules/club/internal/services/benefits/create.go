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
	entity := request.ToMapEntity()
	if err := entity.IsValid(); err != nil {
		log.Printf("ToMapEntity.IsValid.services.benefits.create: %v", err.Error())
		return results.Failure[responses.CreateBenefitResponseDto](ErrInvalidEntity)
	}

	if err := s.preValidation(ctx, request); err != nil {
		return results.Failure[responses.CreateBenefitResponseDto](*err)
	}

	if err := s.benefitRepository.Create(ctx, entity); err != nil {
		if utils.IsTimeout(ctx) {
			return results.Failure[responses.CreateBenefitResponseDto](ErrTimeoutOrCanceled)
		}
		return results.Failure[responses.CreateBenefitResponseDto](ErrCreateBenefit)
	}

	return results.Success(responses.CreateBenefitResponseDto{
		ID: entity.ID,
	})
}

func (s *BenefitService) preValidation(ctx context.Context, request requests.CreateBenefitRequestDto) *results.Error {
	exists, err := s.partnerRepository.ExistsById(ctx, *request.PartnerID)
	if err != nil || !exists {
		if err != nil {
			return &ErrInternalServer
		}

		return &ErrPartnerNotFound
	}

	exists, err = s.contractorRepository.ExistsById(ctx, *request.ContractorID)
	if err != nil || !exists {
		if err != nil {
			return &ErrInternalServer
		}

		return &ErrContractorNotFound
	}

	return nil
}
