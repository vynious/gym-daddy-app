// db/client.go

package db

import (
	"fmt"
	"gd-booking-ms/types"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB      *gorm.DB
	timeout time.Duration
}

func LoadDatabaseConfig() types.GormConfig {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	if host == "" || username == "" || password == "" || dbname == "" {
		log.Fatalf("missing environment variables for database configuration")
	}
	return types.GormConfig{
		Host:     types.Localhost(host),
		Username: types.DatabaseUsername(username),
		Password: types.DatabasePassword(password),
		DBName:   types.DatabaseName(dbname),
	}
}
// Here you would define methods to interact with the Booking data.
// For example, a method to create a booking entry:

type BookingEntry struct {
	ID        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    string
	ClassID   string
	CreatedAt time.Time
}

func SpawnRepository(cfg types.GormConfig) (*Repository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, strconv.Itoa(5432))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	log.Println("Connected to the database successfully.")

	if err := db.AutoMigrate(&BookingEntry{}); err != nil {
		return nil, fmt.Errorf("failed to migrate booking models: %w", err)
	}

	return &Repository{
		DB:      db,
		timeout: time.Duration(2) * time.Second,
	}, nil
}


// CreateBooking creates a new booking record in the database.
func (r *Repository) CreateBooking(userID, classID string) (BookingEntry, error) {
	booking := &BookingEntry{
		UserID:    userID,
		ClassID:   classID,
		CreatedAt: time.Now(),
	}

	err := r.DB.Create(booking).Error
	if err != nil {
		return BookingEntry{}, fmt.Errorf("failed to create booking: %w", err)
	}

	return BookingEntry{
		ID:        booking.ID,
		UserID:    booking.UserID,
		ClassID:   booking.ClassID,
		CreatedAt: booking.CreatedAt,
	}, nil
}

// CloseConnection closes the database connection.
func (r *Repository) CloseConnection() error {
	sqlDB, err := r.DB.DB()
	if err != nil {
		return fmt.Errorf("unable to get sqlDB for closing connections: %w", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("unable to close db: %w", err)
	}
	return nil
}
