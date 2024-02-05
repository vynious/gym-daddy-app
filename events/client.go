package events

import "github.com/segmentio/kafka-go"

type NotificationSubscriber struct {
	subscriber *kafka.Reader
	topic      string
}
