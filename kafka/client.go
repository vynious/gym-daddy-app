package kafka

import "github.com/segmentio/kafka-go"

type NotificationSubscriber struct {
	*kafka.Reader
}

func SpawnNotificationSubscriber(cfg kafka.ReaderConfig) *NotificationSubscriber {
	sub := kafka.NewReader(cfg)
	return &NotificationSubscriber{
		sub,
	}
}
