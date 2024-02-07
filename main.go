package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vynious/gd-telemessenger-ms/bot"
	"github.com/vynious/gd-telemessenger-ms/kafka"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env files")
	}

	// init telegram bot
	telegramUri := os.Getenv("TELEGRAM_BOT_URI")
	if telegramUri == "" {
		log.Fatalf("missing uri for telegram bot")
	}
	telegramBot, err := bot.SpawnBot(telegramUri)
	if err != nil {
		log.Fatalf("")
	}

	telegramBot.PollForUpdates()

	// init kafka subscriber
	kafkaUrl := os.Getenv("KAFKA_URL")
	if kafkaUrl == "" {
		log.Fatalf("missing url for kafka subscriber")
	}
	sub := kafka.SpawnNotificationSubscriber(kafka.LoadKafkaConfigurations())

	defer sub.Close()

	fmt.Println("actively consuming kafka messages...")

	for {
		msg, err := sub.ReadMessage(context.Background()) // msg.Value will be protobufs?
		if err != nil {
			log.Fatalf("error reading kafka messages: %s", err)
		}
		/*
			how does the bot knows which client to send to?
			- needs to be identified through chatId (KV-store of userId and chatId)

			how to get KV-store of userId and chatId?
			-
		*/
	}
}
