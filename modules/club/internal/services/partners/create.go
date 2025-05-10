package partners

import (
	"context"
	"log"

	requests "app/modules/club/domain/dto/requests/partners"
	responses "app/modules/club/domain/dto/responses/partners"
	"app/modules/club/domain/entities"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
)

func (s *PartnerService) Create(ctx context.Context, req requests.CreatePartnerRequestDto) results.Result[responses.CreatePartnerResponseDto] {
	entity := req.ToMapEntity()
	if entity == nil {
		log.Printf("ToMapEntity services.partners.create: %v", ErrInvalidEntity)
		return results.Failure[responses.CreatePartnerResponseDto](ErrInvalidEntity)
	}

	if err := s.validateDuplicates(ctx, entity); err != nil {
		return results.Failure[responses.CreatePartnerResponseDto](*err)
	}

	err := s.partnerRepository.Create(ctx, entity)
	if err != nil {
		if utils.IsTimeout(ctx) {
			return results.Failure[responses.CreatePartnerResponseDto](ErrTimeoutOrCanceled)
		}
		return results.Failure[responses.CreatePartnerResponseDto](ErrCreatePartner)
	}

	return results.Success(
		responses.CreatePartnerResponseDto{
			ID: entity.ID,
		},
	)
}

func (s *PartnerService) validateDuplicates(ctx context.Context, entity *entities.PartnerEntity) *results.Error {
	exists, err := s.partnerRepository.ExistsByCpfCnpj(ctx, entity.CpfCnpj)
	if err != nil {
		return &ErrInternalServer
	}

	if exists {
		return &ErrCpfCnpjDuplicated
	}

	exists, err = s.partnerRepository.ExistsByEmail(ctx, entity.Email)
	if err != nil {
		return &ErrInternalServer
	}

	if exists {
		return &ErrEmailDuplicated
	}

	exists, err = s.partnerRepository.ExistsBySlug(ctx, entity.Slug)
	if err != nil {
		return &ErrInternalServer
	}

	if exists {
		return &ErrSlugDuplicated
	}

	return nil
}
