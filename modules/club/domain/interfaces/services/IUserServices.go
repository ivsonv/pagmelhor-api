package services

import (
	responsesUsers "app/modules/club/domain/dto/responses/users"
	"app/modules/club/domain/results"
	"context"
)

type IUserServices interface {
	Get(ctx context.Context) results.Result[[]responsesUsers.GetUserResponseDto]
}
