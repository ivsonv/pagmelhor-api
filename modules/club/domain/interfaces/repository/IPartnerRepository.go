package repository

import (
	"context"

	"app/modules/club/domain/entities"
)

type IPartnerRepository interface {
	Create(ctx context.Context, entity *entities.PartnerEntity) error
	GetByCpfCnpj(ctx context.Context, cpfCnpj string) (*entities.PartnerEntity, error)
	GetByEmail(ctx context.Context, email string) (*entities.PartnerEntity, error)
	GetBySlug(ctx context.Context, slug string) (*entities.PartnerEntity, error)
}
