package repositories

import (
	"app/modules/club/domain/interfaces/databases"
)

type Repository struct {
	db databases.Database
}

func NewRepository(db databases.Database) *Repository {
	return &Repository{db: db}
}
