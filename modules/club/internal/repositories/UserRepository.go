package repositories

import (
	"context"
	"log"

	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/repository"
)

type UserRepository struct {
	repository *Repository
}

func NewUserRepository(repository *Repository) repository.IUserRepository {
	return &UserRepository{repository: repository}
}

func (r *UserRepository) Get(ctx context.Context) ([]*entities.UserEntity, error) {
	conn := r.repository.db.GetConnection(ctx)
	q := conn.Table(entities.UserEntity{}.TableName())

	users := []*entities.UserEntity{}
	if err := q.Find(&users).Error; err != nil {
		log.Printf("failed to get users: %v\nquery: %s", err, q.Statement.SQL.String())
		return nil, err
	}
	return users, nil
}
