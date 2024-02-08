package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vynious/gd-telemessenger-ms/bot"
	"github.com/vynious/gd-telemessenger-ms/db"
	"github.com/vynious/gd-telemessenger-ms/kafka"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env files")
	}

	// init database
	cfg, err := db.LoadMongoConfig()
	if err == nil {
		log.Fatalf(err.Error())
	}
	repository, err := db.SpawnMongoClient(cfg)
	if err != nil {
		log.Fatalf("failed to start repository")
	}

	// init kafka subscriber
	kafkaUrl := os.Getenv("KAFKA_URL")
	if kafkaUrl == "" {
		log.Fatalf("missing url for kafka subscriber")
	}
	sub := kafka.SpawnNotificationSubscriber(kafka.LoadKafkaConfigurations(), repository)

	// init telegram bot
	telegramUri := os.Getenv("TELEGRAM_BOT_URI")
	if telegramUri == "" {
		log.Fatalf("missing uri for telegram bot")
	}
	telegramBot, err := bot.SpawnBot(telegramUri)
	if err != nil {
		log.Fatalf("")
	}

	defer sub.Subscriber.Close()

	fmt.Println("[kafka] consuming messages...")
	fmt.Println("[telegram-bot] server running...")

	telegramBot.Start()
	sub.Start()

}
