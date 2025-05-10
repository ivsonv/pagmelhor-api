package services

import (
	requests "app/modules/club/domain/dto/requests/benefits"
	responses "app/modules/club/domain/dto/responses/benefits"
	"app/modules/club/domain/results"
	"context"
)

type IBenefitServices interface {
	Create(ctx context.Context, request requests.CreateBenefitRequestDto) results.Result[responses.CreateBenefitResponseDto]
}
