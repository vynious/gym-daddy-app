package events

import "github.com/segmentio/kafka-go"

type NotificationProducer struct {
	producer *kafka.Writer
	topic    string
}

func NewProducer(w *kafka.Writer, topic string) *NotificationProducer {
	return &NotificationProducer{
		producer: w,
		topic:    topic,
	}
}
