package repository

import (
	"context"

	"app/modules/club/domain/entities"
)

type IContractorRepository interface {
	Create(ctx context.Context, contractor *entities.ContractorEntity) error
}
