package repository

import (
	"app/modules/club/domain/entities"
	"context"
)

type IHealthzRepository interface {
	Get(ctx context.Context) (*entities.HealthCheckRepository, error)
	Ping(ctx context.Context) (bool, error)
}
