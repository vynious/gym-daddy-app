package kafka

import (
	"github.com/segmentio/kafka-go"
	"os"
)

func LoadKafkaConfigurations() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers: []string{
			os.Getenv("KAFKA_URL"),
		},
		Topic: "notification",
	}
}
