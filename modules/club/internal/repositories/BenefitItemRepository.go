package repositories

import (
	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type BenefitItemRepository struct {
	repository *Repository
}

func NewBenefitItemRepository(repository *Repository) repository.IBenefitItemRepository {
	return &BenefitItemRepository{
		repository: repository,
	}
}

func (r *BenefitItemRepository) Create(ctx context.Context, benefitItem *entities.BenefitItemEntity) error {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return err
	}

	if err := tx.Create(benefitItem).Error; err != nil {
		log.Printf("failed to repository.benefitItem.create: %v", err)
		return err
	}
	return nil
}

func (r *BenefitItemRepository) GetByID(ctx context.Context, id int) (*entities.BenefitItemEntity, error) {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	benefitItem := entities.BenefitItemEntity{}
	if err := tx.First(&benefitItem, id).Error; err != nil {
		log.Printf("failed to repository.benefitItem.getById: %v", err)
		return nil, err
	}
	return &benefitItem, nil
}

func (r *BenefitItemRepository) GetByBenefitID(ctx context.Context, benefitID int) ([]*entities.BenefitItemEntity, error) {
	tx, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	benefitItems := []*entities.BenefitItemEntity{}
	if err := tx.Where("benefit_id = ?", benefitID).Find(&benefitItems).Error; err != nil {
		log.Printf("failed to repository.benefitItem.getByBenefitId: %v", err)
		return nil, err
	}
	return benefitItems, nil
}

func (r *BenefitItemRepository) getTransaction(ctx context.Context, includeDeleted bool) (*gorm.DB, error) {
	conn, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		log.Printf("failed to get getTransaction repository.benefitItem.getTransaction %v", err)
		return nil, err
	}

	q := conn.Table(entities.BenefitItemEntity{}.TableName())

	if !includeDeleted {
		q = q.Where("deleted_at IS NULL")
	}
	return q, nil
}
