package services

import (
	responsesUsers "app/modules/club/domain/dto/responses/users"
	"context"
)

type IUserServices interface {
	Get(ctx context.Context) ([]responsesUsers.GetUserResponseDto, error)
}
