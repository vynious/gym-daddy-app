package db

import (
	"fmt"
	"github.com/vynious/gym-daddy/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// notifications to be stored in nosql

type Repository struct {
	gc      *gorm.DB
	timeout time.Duration
}

func SpawnRepository(cfg types.GormConfig) (*Repository, error) {

	// Format the DSN string based on the configuration
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, 5432)

	// Open the DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Repository{
		gc:      db,
		timeout: time.Duration(2) * time.Second,
	}, nil
}

func (r *Repository) Close() error {
	sqlDB, err := r.gc.DB()
	if err != nil {
		return fmt.Errorf("unable to get sqlDB for closing connections")
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("unable to close db: %w", err)
	}
}
