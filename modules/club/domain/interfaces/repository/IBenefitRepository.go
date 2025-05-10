package repository

import (
	"app/modules/club/domain/entities"
	"context"
)

type IBenefitRepository interface {
	Create(ctx context.Context, benefit *entities.BenefitEntity) error
	GetByID(ctx context.Context, id int) (*entities.BenefitEntity, error)
	GetByContractorID(ctx context.Context, contractorID int) ([]*entities.BenefitEntity, error)
	GetByPartnerID(ctx context.Context, partnerID int) ([]*entities.BenefitEntity, error)
}
