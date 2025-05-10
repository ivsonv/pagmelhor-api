package services

import (
	"context"

	requests "app/modules/club/domain/dto/requests/contractors"
	responses "app/modules/club/domain/dto/responses/contractors"
	"app/modules/club/domain/results"
)

type IContractorServices interface {
	Create(ctx context.Context, request requests.ContractorRequestDto) results.Result[responses.ContractorResponseDto]
}
