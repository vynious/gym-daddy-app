package kafka

import (
	"github.com/vynious/gym-daddy/types"
	"log"
	"os"
)

func LoadKafkaConfigurations() types.KafkaWriterConfig {
	kafkaUrl := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")
	if kafkaUrl == "" || topic == "" {
		log.Fatalf("missing environment variables")
	}
	return types.KafkaWriterConfig{
		Url:   types.KafkaURL(kafkaUrl),
		Topic: types.KafkaTopic(topic),
	}
}
