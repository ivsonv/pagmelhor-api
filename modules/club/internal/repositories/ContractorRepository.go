package repositories

import (
	"context"
	"log"

	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"

	"gorm.io/gorm"
)

type ContractorRepository struct {
	repository *Repository
}

func NewContractorRepository(repository *Repository) repository.IContractorRepository {
	return &ContractorRepository{repository: repository}
}

func (r *ContractorRepository) Create(ctx context.Context, contractor *entities.ContractorEntity) error {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return err
	}

	if err := q.Create(&contractor).Error; err != nil {
		log.Printf("failed to repository.contractor.create: %v\nquery: %s", err, q.Statement.SQL.String())
		return err
	}

	return nil
}

func (r *ContractorRepository) GetByCpfCnpj(ctx context.Context, cpfCnpj string) (*entities.ContractorEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	contractor := entities.ContractorEntity{}
	if err := q.Where("cpf_cnpj = ?", cpfCnpj).First(&contractor).Error; err != nil {
		log.Printf("failed to repository.contractor.getByCpfCnpj: %v", err)
		return nil, err
	}

	return &contractor, nil
}

func (r *ContractorRepository) GetByEmail(ctx context.Context, email string) (*entities.ContractorEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	contractor := entities.ContractorEntity{}
	if err := q.Where("email = ?", email).First(&contractor).Error; err != nil {
		log.Printf("failed to repository.contractor.getByEmail: %v", err)
		return nil, err
	}

	return &contractor, nil
}

func (r *ContractorRepository) GetBySlug(ctx context.Context, slug string) (*entities.ContractorEntity, error) {
	q, err := r.getTransaction(ctx, false)
	if err != nil {
		return nil, err
	}

	contractor := entities.ContractorEntity{}
	if err := q.Where("slug = ?", slug).First(&contractor).Error; err != nil {
		log.Printf("failed to repository.contractor.getBySlug: %v", err)
		return nil, err
	}

	return &contractor, nil
}

func (r *ContractorRepository) getTransaction(ctx context.Context, includeDeleted bool) (*gorm.DB, error) {
	conn, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		log.Printf("failed to get getTransaction repository.contractor.getTransaction %v", err)
		return nil, err
	}

	q := conn.Table(entities.ContractorEntity{}.TableName())

	// Se includeDeleted for true, busca todos os registros
	if !includeDeleted {
		q = q.Where("deleted_at IS NULL")
	}
	return q, nil
}
