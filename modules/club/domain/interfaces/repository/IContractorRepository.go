package repository

import (
	"context"

	"app/modules/club/domain/entities"
)

type IContractorRepository interface {
	Create(ctx context.Context, entity *entities.ContractorEntity) error
	GetByCpfCnpj(ctx context.Context, cpfCnpj string) (*entities.ContractorEntity, error)
	GetByEmail(ctx context.Context, email string) (*entities.ContractorEntity, error)
	GetBySlug(ctx context.Context, slug string) (*entities.ContractorEntity, error)
}
