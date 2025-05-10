package repositories

import (
	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type BenefitRepository struct {
	repository *Repository
}

func NewBenefitRepository(repository *Repository) repository.IBenefitRepository {
	return &BenefitRepository{
		repository: repository,
	}
}

func (r *BenefitRepository) Create(ctx context.Context, benefit *entities.BenefitEntity) error {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return err
	}

	if err := tx.Create(benefit).Error; err != nil {
		log.Printf("failed to repository.benefit.create: %v", err)
		return err
	}
	return nil
}

func (r *BenefitRepository) GetByID(ctx context.Context, id int) (*entities.BenefitEntity, error) {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	benefit := entities.BenefitEntity{}
	if err := tx.First(&benefit, id).Error; err != nil {
		log.Printf("failed to repository.benefit.getById: %v", err)
		return nil, err
	}
	return &benefit, nil
}

func (r *BenefitRepository) GetByContractorID(ctx context.Context, contractorID int) ([]*entities.BenefitEntity, error) {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	benefits := []*entities.BenefitEntity{}
	if err := tx.Where("contractor_id = ?", contractorID).Find(&benefits).Error; err != nil {
		log.Printf("failed to repository.benefit.getByContractorId: %v", err)
		return nil, err
	}
	return benefits, nil
}

func (r *BenefitRepository) GetByPartnerID(ctx context.Context, partnerID int) ([]*entities.BenefitEntity, error) {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	benefits := []*entities.BenefitEntity{}
	if err := tx.Where("partner_id = ?", partnerID).Find(&benefits).Error; err != nil {
		log.Printf("failed to repository.benefit.getByPartnerId: %v", err)
		return nil, err
	}
	return benefits, nil
}

func (r *BenefitRepository) getTransaction(ctx context.Context, includeDeleted bool) (*gorm.DB, error) {
	conn, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		log.Printf("failed to get getTransaction repository.benefit.getTransaction %v", err)
		return nil, err
	}

	q := conn.Table(entities.BenefitEntity{}.TableName())

	if !includeDeleted {
		q = q.Where("deleted_at IS NULL")
	}
	return q, nil
}
