package repositories

import (
	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"
	"context"
	"log"
	"os"
)

type HealthzRepository struct {
	repository *Repository
}

func NewHealthzRepository(repository *Repository) repository.IHealthzRepository {
	return &HealthzRepository{repository: repository}
}

func (r *HealthzRepository) Get(ctx context.Context) (*entities.HealthCheckRepository, error) {
	db, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		log.Printf("failed to get database connection in repository.healthz.get: %v", err)
		return nil, err
	}

	openConnections := 0
	err = db.Raw("SELECT count(*)::int FROM pg_stat_activity WHERE datname = ?", os.Getenv("DB_NAME")).Scan(&openConnections).Error
	if err != nil {
		log.Printf("failed to repository.healthz.get: %v", err)
		return nil, err
	}

	version := ""
	err = db.Raw("SHOW server_version").Scan(&version).Error
	if err != nil {
		log.Printf("failed to repository.healthz.get: %v", err)
		return nil, err
	}

	maxConnections := 0
	err = db.Raw("SHOW max_connections").Scan(&maxConnections).Error
	if err != nil {
		log.Printf("failed to repository.healthz.get: %v", err)
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
	db, err := r.repository.db.GetConnection(ctx)
	if err != nil {
		return false, err
	}

	var rows []any
	err = db.Raw("SELECT 1").Scan(&rows).Error
	if err != nil {
		log.Printf("failed to repository.healthz.ping: %v", err)
		return false, err
	}

	return len(rows) > 0, nil
}
