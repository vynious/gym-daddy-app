package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

type (
	DBConnector struct {
		DB      *pg.DB
		Options *DBConnectorOptions
	}

	DBConnectorOptions struct {
		User     string
		DB_Name  string
		Addr     string
		Password string
	}
)

func InitDBConnection() *DBConnector {
	dbConnector := &DBConnector{
		Options: &DBConnectorOptions{
			User:     os.Getenv("DB_USER"),
			DB_Name:  os.Getenv("DB_NAME"),
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	maxRetries := 5
	retryDelay := 5 * time.Second

	var db *pg.DB
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Connecting to database (attempt %d of %d)...\n", attempt, maxRetries)
		pgOptions := &pg.Options{
			User:     dbConnector.Options.User,
			Database: dbConnector.Options.DB_Name,
			Addr:     dbConnector.Options.Addr,
			Password: dbConnector.Options.Password,
		}
		db = pg.Connect(pgOptions)
		_, err = db.Exec("SELECT 1")
		if err == nil {
			fmt.Println("Successfully connected to the database.")
			dbConnector.DB = db
			return dbConnector
		}

		fmt.Printf("Failed to connect to the database: %v\n", err)
		if attempt < maxRetries {
			fmt.Printf("Retrying in %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	log.Fatalf("Failed to connect to the database after %d attempts: %v", maxRetries, err)
	return nil
}
