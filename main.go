package main

import (
	"github.com/vynious/gd-telemessenger-ms/bot"
	"log"
	"os"
)

func main() {

	uri := os.Getenv("TELEGRAM_BOT_URI")
	if uri == "" {
		log.Fatalf("missing uri for telegram bot")
	}
	botSvc, err := bot.SpawnBot(uri)
	if err != nil {
		log.Fatalf("")
	}
	botSvc.Client
}
