package repository

import (
	"context"

	"app/modules/club/domain/entities"
)

type IUserRepository interface {
	Get(ctx context.Context) ([]*entities.UserEntity, error)
}
