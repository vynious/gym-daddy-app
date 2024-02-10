package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/vynious/gym-daddy/db"
	"github.com/vynious/gym-daddy/types"
	"log"
)

type NotificationProducer struct {
	Producer *kafka.Writer
	Database *db.Repository
}

func SpawnKafkaProducer(cfg types.KafkaWriterConfig, repo *db.Repository) *NotificationProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(string(cfg.Url)),
		Topic:    string(cfg.Topic),
		Balancer: &kafka.LeastBytes{},
	}
	return &NotificationProducer{
		Producer: w,
		Database: repo,
	}
}

func (np *NotificationProducer) CloseConnection() {
	if err := np.Producer.Close(); err != nil {
		log.Fatalf(err.Error())
	}
	if err := np.Database.Close(); err != nil {
		log.Fatalf(err.Error())
	}
}
