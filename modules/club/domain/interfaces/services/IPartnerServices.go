package services

import (
	"context"

	requests "app/modules/club/domain/dto/requests/partners"
	responses "app/modules/club/domain/dto/responses/partners"
	"app/modules/club/domain/results"
)

type IPartnerServices interface {
	Create(ctx context.Context, request requests.CreatePartnerRequestDto) results.Result[responses.CreatePartnerResponseDto]
}
