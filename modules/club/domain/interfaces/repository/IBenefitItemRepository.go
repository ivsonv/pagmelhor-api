package repository

import (
	"app/modules/club/domain/entities"
	"context"
)

type IBenefitItemRepository interface {
	Create(ctx context.Context, benefitItem *entities.BenefitItemEntity) error
	GetByID(ctx context.Context, id int) (*entities.BenefitItemEntity, error)
	GetByBenefitID(ctx context.Context, benefitID int) ([]*entities.BenefitItemEntity, error)
}
