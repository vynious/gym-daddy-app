package db

import (
	"github.com/vynious/gym-daddy/types"
	"log"
	"os"
)

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
