package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vynious/gd-telemessenger-ms/bot"
	"github.com/vynious/gd-telemessenger-ms/db"
	"github.com/vynious/gd-telemessenger-ms/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// init telegram bot

	telegramBot, err := bot.SpawnBot(bot.LoadBotConfig(), repository)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// init kafka subscriber

	sub := kafka.SpawnNotificationSubscriber(kafka.LoadKafkaConfigurations(), telegramBot)

	//defer sub.CloseConnections() // close both mongo and kafka clients

	fmt.Println("[kafka] consuming messages...")
	fmt.Println("[telegram-bot] server running...")

	go telegramBot.Start()
	go sub.Start()

	// Block main goroutine until an OS signal is received
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	fmt.Println("Termination signal received, shutting down.")

	// Clean up resources
	sub.CloseConnections()

}
