package rpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vynious/gym-daddy/db"
	"github.com/vynious/gym-daddy/kafka"
	"github.com/vynious/gym-daddy/pb/proto_files/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"sync"
)

type Client struct {
	Repository *db.Repository
	GConn      *notification.NotificationServiceClient
	Notifier   *kafka.NotificationProducer
	*notification.UnimplementedNotificationServiceServer
}

func SpawnGrpcClient(repo *db.Repository, notifier *kafka.NotificationProducer) (*Client, error) {
	grpcServerUrl := os.Getenv("NOTIFICATION_SERVER_URL")
	if grpcServerUrl == "" {
		return nil, fmt.Errorf("check grpc server environment variable")
	}

	clientConn, err := grpc.Dial(grpcServerUrl, grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("error calling grpc server: %w", err)
	}
	notificationClient := notification.NewNotificationServiceClient(clientConn)
	return &Client{
		Repository: repo,
		GConn:      &notificationClient,
		Notifier:   notifier,
	}, nil
}

func (c *Client) CreateNotification(ctx context.Context, req *notification.CreateNotificationRequest) (*notification.CreateNotificationResponse, error) {

	notificationProto := req.GetNotification()

	if len(notificationProto.Id) > 0 {

	} else {
		id, err := uuid.NewUUID()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate uuid for notification: %v", err)
		}
		notificationProto.Id = id.String()
	}

	errCh := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.Repository.SaveNotification(notificationProto); err != nil {
			errCh <- status.Errorf(codes.Internal, "failed to save notitfication to db: %w", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.Notifier.SendNotification(ctx, notificationProto); err != nil {
			errCh <- status.Errorf(codes.Internal, "failed to send notitfication to kafka queue: %w", err)
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		return nil, err
	}

	return &notification.CreateNotificationResponse{}, nil
}

func (c *Client) CloseConnections() {
	var errors []error
	if err := c.Notifier.CloseConnection(); err != nil {
		errors = append(errors, err)
	}
	if err := c.Repository.CloseConnection(); err != nil {
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		log.Fatalf("failed to close connections %v", errors)
	}
}
