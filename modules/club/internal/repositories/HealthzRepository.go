package repositories

import (
	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"
	"context"
	"os"
)

type HealthzRepository struct {
	repository *Repository
}

func NewHealthzRepository(repository *Repository) repository.IHealthzRepository {
	return &HealthzRepository{repository: repository}
}

func (r *HealthzRepository) Get(ctx context.Context) (*entities.HealthCheckRepository, error) {
	db := r.repository.db.GetConnection(ctx)

	var openConnections int
	err := db.Raw("SELECT count(*)::int FROM pg_stat_activity WHERE datname = ?", os.Getenv("DB_NAME")).Scan(&openConnections).Error
	if err != nil {
		return nil, err
	}

	var version string
	err = db.Raw("SHOW server_version").Scan(&version).Error
	if err != nil {
		return nil, err
	}

	var maxConnections int
	err = db.Raw("SHOW max_connections").Scan(&maxConnections).Error
	if err != nil {
		return nil, err
	}

	return &entities.HealthCheckRepository{
		AvailableConnections: maxConnections - openConnections,
		OpenConnections:      openConnections,
		MaxConnections:       maxConnections,
		Version:              version,
	}, nil
}

func (r *HealthzRepository) Ping(ctx context.Context) (bool, error) {
	db := r.repository.db.GetConnection(ctx)

	var rows []any
	err := db.Raw("SELECT 1").Scan(&rows).Error
	if err != nil {
		return false, err
	}

	return len(rows) > 0, nil
}
