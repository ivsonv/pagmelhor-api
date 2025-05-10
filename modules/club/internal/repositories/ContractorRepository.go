package repositories

import (
	"context"
	"log"

	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"
)

type ContractorRepository struct {
	repository *Repository
}

func NewContractorRepository(repository *Repository) repository.IContractorRepository {
	return &ContractorRepository{repository: repository}
}

func (r *ContractorRepository) Create(ctx context.Context, contractor *entities.ContractorEntity) error {
	conn := r.repository.db.GetConnection(ctx)
	q := conn.Table(entities.ContractorEntity{}.TableName())

	if err := q.Create(&contractor).Error; err != nil {
		log.Printf("failed to create contractor: %v\nquery: %s", err, q.Statement.SQL.String())
		return err
	}

	return nil
}
