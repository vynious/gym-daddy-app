package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/vynious/gd-telemessenger-ms/bot"
	"github.com/vynious/gd-telemessenger-ms/types"
	"log"
	"os"
)

type NotificationSubscriber struct {
	Subscriber  *kafka.Reader
	TelegramBot *bot.Bot
}

func LoadKafkaConfigurations() kafka.ReaderConfig {
	kafkaUrl := os.Getenv("KAFKA_URL")
	if kafkaUrl == "" {
		log.Fatalf("missing url for kafka subscriber")
	}

	return kafka.ReaderConfig{
		Brokers: []string{
			kafkaUrl,
		},
		Topic: "notification",
	}
}

func SpawnNotificationSubscriber(cfg kafka.ReaderConfig, b *bot.Bot) *NotificationSubscriber {
	sub := kafka.NewReader(cfg)
	return &NotificationSubscriber{
		Subscriber:  sub,
		TelegramBot: b,
	}
}

func (ns *NotificationSubscriber) Start() {

	for {
		msg, err := ns.Subscriber.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error reading kafka messages: %v", err)

		}
		var nm types.NotificationMessage
		if err := json.Unmarshal(msg.Value, &nm); err != nil {
			log.Println("failed to unmarshal msg: ", err)
			continue
		}
		log.Printf("%v", nm)
		chatId, err := ns.TelegramBot.Database.GetSubscription(nm.TelegramHandle)
		if err != nil {
			log.Println("failed to get chatID: ", err)
			continue
		}

		if err := ns.TelegramBot.SendNotification(chatId, &nm); err != nil {
			log.Println("failed to send notification to user: ", err)
			continue
		}
	}
}

func (ns *NotificationSubscriber) CloseConnections() {
	if err := ns.Subscriber.Close(); err != nil {
		log.Fatalf("error closing kafka connection: %v", err)
	}
	if err := ns.TelegramBot.Database.CloseConnection(); err != nil {
		log.Fatalf("error closing mongodb connection: %v", err)
	}
}
