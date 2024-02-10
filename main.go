package main

import (
	"github.com/joho/godotenv"
	"github.com/vynious/gym-daddy/db"
	"github.com/vynious/gym-daddy/kafka"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env files")
	}

	// init database
	repository, err := db.SpawnRepository(db.LoadDatabaseConfig())
	if err != nil {
		log.Fatalf("failed to start repository")
	}

	// init kafka producer
	prod := kafka.SpawnKafkaProducer(kafka.LoadKafkaConfigurations(), repository)

	defer prod.CloseConnection()
}
