package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/vynious/gym-daddy/gd-notification-ms/pb/proto_files/notification"
	"github.com/vynious/gym-daddy/gd-notification-ms/types"
	"log"
	"os"
)

type NotificationProducer struct {
	producer *kafka.Writer
}

type NotificationMessage struct {
	TelegramHandle string
	EventType      string
	Content        string
}

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

func SpawnKafkaProducer(cfg types.KafkaWriterConfig) *NotificationProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(string(cfg.Url)),
		Topic:    string(cfg.Topic),
		Balancer: &kafka.LeastBytes{},
	}
	return &NotificationProducer{
		producer: w,
	}
}

func (np *NotificationProducer) SendNotification(ctx context.Context, notification *notification.Notification) error {
	var n NotificationMessage
	n.TelegramHandle = notification.GetTelegramHandle()
	n.Content = notification.GetContent()
	n.EventType = notification.GetNotificationType()

	jsonMsg, err := json.Marshal(n)
	if err != nil {
		return fmt.Errorf("failed to marshal notification event: %w", err)
	}
	if err := np.producer.WriteMessages(ctx, kafka.Message{Value: jsonMsg}); err != nil {
		return fmt.Errorf("failed to write message to kafka queue :%w", err)
	}
	return nil
}

func (np *NotificationProducer) CloseConnection() error {
	if err := np.producer.Close(); err != nil {
		return fmt.Errorf("failed to close kafka connection %w", err)
	}
	return nil
}
