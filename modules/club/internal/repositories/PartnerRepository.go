package repositories

import (
	"context"
	"log"

	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"

	"gorm.io/gorm"
)

type PartnerRepository struct {
	repository *Repository
}

func NewPartnerRepository(repository *Repository) repository.IPartnerRepository {
	return &PartnerRepository{repository: repository}
}

func (r *PartnerRepository) Create(ctx context.Context, partner *entities.PartnerEntity) error {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return err
	}

	if err := q.Create(&partner).Error; err != nil {
		log.Printf("failed to repository.partner.create: %v\nquery: %s", err, q.Statement.SQL.String())
		return err
	}

	return nil
}

func (r *PartnerRepository) GetByCpfCnpj(ctx context.Context, cpfCnpj string) (*entities.PartnerEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	partner := entities.PartnerEntity{}
	if err := q.Where("cpf_cnpj = ?", cpfCnpj).First(&partner).Error; err != nil {
		log.Printf("failed to repository.partner.getByCpfCnpj: %v", err)
		return nil, err
	}

	return &partner, nil
}

func (r *PartnerRepository) GetByEmail(ctx context.Context, email string) (*entities.PartnerEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	partner := entities.PartnerEntity{}
	if err := q.Where("email = ?", email).First(&partner).Error; err != nil {
		log.Printf("failed to repository.partner.getByEmail: %v", err)
		return nil, err
	}

	return &partner, nil
}

func (r *PartnerRepository) GetBySlug(ctx context.Context, slug string) (*entities.PartnerEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	partner := entities.PartnerEntity{}
	if err := q.Where("slug = ?", slug).First(&partner).Error; err != nil {
		log.Printf("failed to repository.partner.getBySlug: %v", err)
		return nil, err
	}

	return &partner, nil
}

func (r *PartnerRepository) GetById(ctx context.Context, id int) (*entities.PartnerEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	partner := entities.PartnerEntity{}
	if err := q.Where("id = ?", id).First(&partner).Error; err != nil {
		log.Printf("failed to repository.partner.getById: %v", err)
		return nil, err
	}

	return &partner, nil
}

func (r *PartnerRepository) ExistsById(ctx context.Context, id int) (bool, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return false, err
	}

	exists := false
	err = q.Select("1").Where("id = ?", id).Find(&exists).Error

	if err != nil {
		log.Printf("failed to repository.partner.existsById: %v", err)
		return false, err
	}

	return exists, nil
}

func (r *PartnerRepository) ExistsByCpfCnpj(ctx context.Context, cpfCnpj string) (bool, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return false, err
	}

	exists := false
	err = q.Select("1").Where("cpf_cnpj = ?", cpfCnpj).Find(&exists).Error

	if err != nil {
		log.Printf("failed to repository.partner.existsByCpfCnpj: %v", err)
		return false, err
	}

	return exists, nil
}

func (r *PartnerRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return false, err
	}

	exists := false
	err = q.Select("1").Where("email = ?", email).Find(&exists).Error

	if err != nil {
		log.Printf("failed to repository.partner.existsByEmail: %v", err)
		return false, err
	}

	return exists, nil
}

func (r *PartnerRepository) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return false, err
	}

	exists := false
	err = q.Select("1").Where("slug = ?", slug).Find(&exists).Error

	if err != nil {
		log.Printf("failed to repository.partner.existsBySlug: %v", err)
		return false, err
	}

	return exists, nil
}

func (r *PartnerRepository) getTransaction(ctx context.Context, includeDeleted bool) (*gorm.DB, error) {
	conn, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		log.Printf("failed to get getTransaction repository.partner.getTransaction %v", err)
		return nil, err
	}

	q := conn.Table(entities.PartnerEntity{}.TableName())

	if !includeDeleted {
		q = q.Where("deleted_at IS NULL")
	}
	return q, nil
}
