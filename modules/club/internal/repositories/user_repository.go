package repositories

import (
	"context"
	"log"

	"app/modules/club/domain/entities"
	"app/modules/club/domain/interfaces/databases"
	"app/modules/club/domain/interfaces/repository"
)

type userRepository struct {
	db databases.Database
}

func NewUserRepository(db databases.Database) repository.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Get(ctx context.Context) ([]*entities.UserEntity, error) {
	conn := r.db.GetConnection(ctx)
	q := conn.Table(entities.UserEntity{}.TableName())

	users := []*entities.UserEntity{}
	if err := q.Find(&users).Error; err != nil {
		log.Printf("failed to get users: %v\nquery: %s", err, q.Statement.SQL.String())
		return nil, err
	}
	return users, nil
}
