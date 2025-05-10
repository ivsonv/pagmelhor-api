package repository

import (
	"context"

	"app/modules/club/domain/entities"
)

type IContractorRepository interface {
	GetByCpfCnpj(ctx context.Context, cpfCnpj string) (*entities.ContractorEntity, error)
	GetByEmail(ctx context.Context, email string) (*entities.ContractorEntity, error)
	GetBySlug(ctx context.Context, slug string) (*entities.ContractorEntity, error)
	GetById(ctx context.Context, id int) (*entities.ContractorEntity, error)
	Create(ctx context.Context, entity *entities.ContractorEntity) error
	ExistsById(ctx context.Context, id int) (bool, error)
}
