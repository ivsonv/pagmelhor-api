package services

import (
	requests "app/modules/club/domain/dto/requests/benefit_items"
	responses "app/modules/club/domain/dto/responses/benefit_items"
	"app/modules/club/domain/results"
	"context"
)

type IBenefitItemServices interface {
	Create(ctx context.Context, request requests.CreateBenefitItemRequestDto) results.Result[responses.CreateBenefitItemResponseDto]
}
