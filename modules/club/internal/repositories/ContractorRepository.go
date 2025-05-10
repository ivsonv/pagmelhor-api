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
	conn, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		log.Printf("failed to get database connection in create contractor: %v", err)
		return err
	}

	q := conn.Table(entities.ContractorEntity{}.TableName())

	if err := q.Create(&contractor).Error; err != nil {
		log.Printf("failed to create contractor: %v\nquery: %s", err, q.Statement.SQL.String())
		return err
	}

	return nil
}
