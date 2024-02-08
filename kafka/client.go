package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/vynious/gd-telemessenger-ms/db"
	"github.com/vynious/gd-telemessenger-ms/types"
	"log"
)

type NotificationSubscriber struct {
	Subscriber *kafka.Reader
	Database   *db.Repository
}

func SpawnNotificationSubscriber(cfg kafka.ReaderConfig, repo *db.Repository) *NotificationSubscriber {
	sub := kafka.NewReader(cfg)
	return &NotificationSubscriber{
		Subscriber: sub,
		Database:   repo,
	}
}

func (ns *NotificationSubscriber) Start() {
	for {
		msg, err := ns.Subscriber.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error reading kafka messages: %s", err)

		}
		telHandle := types.TelegramHandle(msg.Key)

		chatId, err := ns.Database.GetSubscription(telHandle)
		if err != nil {

		}

		/*
			how does the bot knows which client to send to?
			- needs to be identified through chatId (KV-store of userId and chatId)

			how to get KV-store of userId and chatId?
			-
		*/
	}
}
