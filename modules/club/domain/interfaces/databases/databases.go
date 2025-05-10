package databases

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Database interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	GetConnection(ctx context.Context) (*gorm.DB, error)
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
}

type Cache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string, data any) error
	Delete(ctx context.Context, key string) error
	Ping(ctx context.Context) error
}
