package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

type LogProducer struct {
	producer *kafka.Writer
}

type KafkaWriterConfig struct {
	Url   string
	Topic string
}

func LoadKafkaConfigurations() KafkaWriterConfig {
	kafkaUrl := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")
	if kafkaUrl == "" || topic == "" {
		log.Fatalf("missing environment variables")
	}
	return KafkaWriterConfig{
		Url:   kafkaUrl,
		Topic: topic,
	}
}

func SpawnKafkaProducer(cfg KafkaWriterConfig) *LogProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Url),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &LogProducer{
		producer: w,
	}
}

func (lp *LogProducer) SendLog(ctx context.Context, contentType string, content string) error {

	log.Println("sending error log")

	var l struct {
		RequestId   string
		Content     string
		ContentType string
	}

	rid := ctx.Value("request_id")
	strRID, ok := rid.(string)
	if !ok {
		log.Println("failed to convert to string")
	}

	l.RequestId = strRID
	l.Content = content
	l.ContentType = contentType

	log.Printf("RequestId: %s, Content: %s", strRID, content)

	jsonMsg, err := json.Marshal(l)

	if err != nil {
		return fmt.Errorf("failed to marshal log event: %w", err)
	}
	if err := lp.producer.WriteMessages(ctx, kafka.Message{Value: jsonMsg}); err != nil {
		return fmt.Errorf("failed to write message to kafka queue :%w", err)
	}

	return nil
}

func (lp *LogProducer) CloseConnection() error {
	if err := lp.producer.Close(); err != nil {
		return fmt.Errorf("failed to close kafka connection %w", err)
	}
	return nil
}
