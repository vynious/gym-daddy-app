package grpc

import (
	"github.com/vynious/gym-daddy/kafka"
)

type Client struct {
	// grpc client
	// mongodb
	notifier *kafka.NotificationProducer
}

func SpawnGrpcClient(notifier *kafka.NotificationProducer) (*Client, error) {
	// create grpc client from the generated client code... (not yet imported)
	return &Client{
		notifier: notifier,
	}, nil
}

func (c *Client) ReceiveMessage() {
	// store message into to database
	// send message to kafka

}
