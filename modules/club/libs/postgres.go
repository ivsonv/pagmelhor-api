package databases

import (
	"app/configs"
	"app/modules/club/domain"
	"app/modules/club/domain/interfaces/databases"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	ConnStr       string
	MaxOpenConns  int
	MaxIdleConns  int
	ConnLifetime  int
	MaxRetries    int
	RetryInterval int
}

type postgresImpl struct {
	PostgresConfig
	conn *gorm.DB
	tx   *gorm.DB
}

func NewPostgres(cfg *configs.Config) databases.Database {
	sslMode := "require"
	if cfg.DBSSLMode != "" {
		sslMode = cfg.DBSSLMode
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=%s TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, domain.SchemaClubName, sslMode)

	return &postgresImpl{
		PostgresConfig: PostgresConfig{
			ConnStr:       connStr,
			MaxOpenConns:  20,
			MaxIdleConns:  2,
			ConnLifetime:  20,
			MaxRetries:    2,
			RetryInterval: 3,
		},
	}
}

func (p *postgresImpl) GetConnection(ctx context.Context) (*gorm.DB, error) {
	if p.conn == nil {
		var err error
		for i := range p.MaxRetries {
			p.conn, err = p.connect(ctx)
			if err == nil {
				return p.conn, nil
			}
			log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, p.MaxRetries, err)
			if i < p.MaxRetries-1 {
				time.Sleep(time.Duration(p.RetryInterval) * time.Second)
			}
		}
		log.Printf("Could not connect to database after %d attempts", p.MaxRetries)
		return nil, nil
	}

	if ctx.Err() != nil {
		expired := ctx.Err() == context.DeadlineExceeded || ctx.Err() == context.Canceled
		if expired {
			log.Printf("Context expired before database operation: %v", ctx.Err())
			return nil, ctx.Err()
		} else {
			log.Printf("Context expired before database operation: %v", ctx.Err())
			return nil, ctx.Err()
		}
	}
	return p.conn, nil
}

func (p *postgresImpl) BeginTransaction(ctx context.Context) (context.Context, error) {
	db, _ := p.GetConnection(ctx)
	if db == nil {
		return ctx, fmt.Errorf("database connection not available")
	}
	tx := db.Begin()
	p.tx = tx
	return ctx, tx.Error
}

func (p *postgresImpl) Rollback(ctx context.Context) (context.Context, error) {
	if p.tx != nil {
		return ctx, p.tx.Rollback().Error
	}
	return ctx, nil
}

func (p *postgresImpl) Commit(ctx context.Context) (context.Context, error) {
	if p.tx != nil {
		return ctx, p.tx.Commit().Error
	}
	return ctx, nil
}

func (p *postgresImpl) Ping(ctx context.Context) error {
	db, err := p.GetConnection(ctx)
	if db == nil {
		log.Printf("database connection not available %v", err)
		return fmt.Errorf("database connection not available")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

func (p *postgresImpl) Close(ctx context.Context) error {
	if p.conn != nil {
		sqlDB, err := p.conn.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

func (db *postgresImpl) connect(_ context.Context) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	var sqlConn *sql.DB
	var err error

	sqlConn, err = sql.Open("postgres", db.ConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlConn.SetMaxOpenConns(db.MaxOpenConns)
	sqlConn.SetMaxIdleConns(db.MaxIdleConns)
	sqlConn.SetConnMaxLifetime(time.Second * time.Duration(db.ConnLifetime))

	gormConn, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlConn}), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gorm: %w", err)
	}

	return gormConn, nil
}
