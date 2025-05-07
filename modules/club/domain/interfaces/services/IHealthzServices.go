package services

import (
	"context"

	responses "app/modules/club/domain/dto/responses/healthz"
	"app/modules/club/domain/results"
)

type IHealthzServices interface {
	Get(ctx context.Context) results.Result[responses.GetHealthzResponseDto]
}
