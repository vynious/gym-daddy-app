package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vynious/gd-telemessenger-ms/bot"
	"github.com/vynious/gd-telemessenger-ms/db"
	"github.com/vynious/gd-telemessenger-ms/kafka"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env files")
	}

	// init database

	repository, err := db.SpawnRepository(db.LoadMongoConfig())
	if err != nil {
		log.Fatalf("failed to start repository")
	}

	// init kafka subscriber

	sub := kafka.SpawnNotificationSubscriber(kafka.LoadKafkaConfigurations(), repository)

	// init telegram bot

	telegramBot, err := bot.SpawnBot(bot.LoadBotConfig(), repository)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer sub.CloseConnections() // close both mongo and kafka clients

	fmt.Println("[kafka] consuming messages...")
	fmt.Println("[telegram-bot] server running...")

	go telegramBot.Start()
	go sub.Start()

}
