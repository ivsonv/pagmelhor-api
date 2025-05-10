package contractors

import (
	"context"
	"log"

	requests "app/modules/club/domain/dto/requests/contractors"
	responses "app/modules/club/domain/dto/responses/contractors"
	"app/modules/club/domain/results"
)

func (s *ContractorService) Create(ctx context.Context, req requests.ContractorRequestDto) results.Result[responses.ContractorResponseDto] {
	// Verifica se o contexto já está expirado
	if ctx.Err() != nil {
		log.Printf("Context expired before operation: %v", ctx.Err())
		return results.Failure[responses.ContractorResponseDto](ErrTimeout)
	}

	entity := req.ToMapEntity()
	if entity == nil {
		log.Printf("ToMapEntity create contractor: %v", ErrInvalidEntity)
		return results.Failure[responses.ContractorResponseDto](ErrInvalidEntity)
	}

	err := s.contractorRepository.Create(ctx, entity)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Create contractor timeout: %v", ctx.Err())
			return results.Failure[responses.ContractorResponseDto](ErrTimeout)
		}
		return results.Failure[responses.ContractorResponseDto](ErrCreateContractor)
	}

	return results.Success(
		responses.ContractorResponseDto{
			ID: entity.ID,
		},
	)
}
