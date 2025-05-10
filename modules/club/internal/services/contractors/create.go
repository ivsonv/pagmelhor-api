package contractors

import (
	"context"
	"log"

	requests "app/modules/club/domain/dto/requests/contractors"
	responses "app/modules/club/domain/dto/responses/contractors"
	"app/modules/club/domain/entities"
	"app/modules/club/domain/results"
	"app/modules/club/utils"
)

func (s *ContractorService) Create(ctx context.Context, req requests.ContractorRequestDto) results.Result[responses.ContractorResponseDto] {
	entity := req.ToMapEntity()
	if entity == nil {
		log.Printf("ToMapEntity services.create contractor: %v", ErrInvalidEntity)
		return results.Failure[responses.ContractorResponseDto](ErrInvalidEntity)
	}

	if err := s.validateDuplicates(ctx, entity); err != nil {
		return results.Failure[responses.ContractorResponseDto](*err)
	}

	err := s.contractorRepository.Create(ctx, entity)
	if err != nil {
		if utils.IsTimeout(ctx) {
			return results.Failure[responses.ContractorResponseDto](ErrTimeoutOrCanceled)
		}

		return results.Failure[responses.ContractorResponseDto](ErrCreateContractor)
	}

	return results.Success(
		responses.ContractorResponseDto{
			ID: entity.ID,
		},
	)
}

func (s *ContractorService) validateDuplicates(ctx context.Context, entity *entities.ContractorEntity) *results.Error {
	contractor, _ := s.contractorRepository.GetByCpfCnpj(ctx, entity.CpfCnpj)
	if contractor != nil {
		return &ErrCpfCnpjDuplicated
	}

	contractor, _ = s.contractorRepository.GetByEmail(ctx, entity.Email)
	if contractor != nil {
		return &ErrEmailDuplicated
	}

	contractor, _ = s.contractorRepository.GetBySlug(ctx, entity.Slug)
	if contractor != nil {
		return &ErrSlugDuplicated
	}

	return nil
}
