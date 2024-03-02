package db

import (
	"fmt"
	"github.com/vynious/gym-daddy/pb/proto_files/notification"
	"github.com/vynious/gym-daddy/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// notifications to be stored in nosql

type Repository struct {
	gc      *gorm.DB
	timeout time.Duration
}

func LoadDatabaseConfig() types.GormConfig {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	if host == "" || username == "" || password == "" || dbname == "" {
		log.Fatalf("missing environment variables")
	}
	return types.GormConfig{
		Host:     types.Localhost(host),
		Username: types.DatabaseUsername(username),
		Password: types.DatabasePassword(password),
		DBName:   types.DatabaseName(dbname),
	}
}

type NotificationEntry struct {
	ID             string    `gorm:"primary key" json:"id"`
	TelegramHandle string    `json:"telegram_handle"`
	EventType      string    `json:"event_type"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

func SpawnRepository(cfg types.GormConfig) (*Repository, error) {
	var db *gorm.DB
	var err error
	maxAttempts := 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai",
			cfg.Host, cfg.Username, cfg.Password, cfg.DBName, 5432)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Database connection attempt %d failed: %v", attempt, err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database after %d attempts", maxAttempts)
	}

	if err := db.AutoMigrate(&NotificationEntry{}); err != nil {
		return nil, fmt.Errorf("failed to migrate models : %w", err)
	}

	return &Repository{
		gc:      db,
		timeout: time.Duration(2) * time.Second,
	}, nil // Connection successful

}

func (r *Repository) SaveNotification(notification *notification.Notification) error {
	var n NotificationEntry

	n.TelegramHandle = notification.GetTelegramHandle()
	n.Content = notification.GetContent()
	n.EventType = notification.GetEventType()
	n.CreatedAt = notification.GetCreatedAt().AsTime()

	err := r.gc.Create(n).Error
	if err != nil {
		return fmt.Errorf("failed to create notification : %w", err)
	}

	return nil
}

func (r *Repository) CloseConnection() error {
	sqlDB, err := r.gc.DB()
	if err != nil {
		return fmt.Errorf("unable to get sqlDB for closing connections")
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("unable to close db: %w", err)
	}
	return nil
}
