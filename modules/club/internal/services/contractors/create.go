package contractors

import (
	"context"

	requests "app/modules/club/domain/dto/requests/contractors"
	responses "app/modules/club/domain/dto/responses/contractors"
	"app/modules/club/domain/results"
)

func (s *ContractorService) Create(ctx context.Context, req requests.ContractorRequestDto) results.Result[responses.ContractorResponseDto] {
	entity := req.ToMapEntity()

	err := s.contractorRepository.Create(ctx, entity)
	if err != nil {
		return results.Failure[responses.ContractorResponseDto](ErrCreateContractor)
	}

	return results.Success(
		responses.ContractorResponseDto{
			ID: entity.ID,
		},
	)
}
