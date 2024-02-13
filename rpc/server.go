package rpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/vynious/gym-daddy/db"
	"github.com/vynious/gym-daddy/kafka"
	"github.com/vynious/gym-daddy/pb/proto_files/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type Server struct {
	Repository *db.Repository
	Notifier   *kafka.NotificationProducer
	*notification.UnimplementedNotificationServiceServer
}

func SpawnGrpcServer(repo *db.Repository, notifier *kafka.NotificationProducer) (*Server, error) {
	return &Server{
		Repository: repo,
		Notifier:   notifier,
	}, nil
}

func (c *Server) CreateNotification(ctx context.Context, req *notification.CreateNotificationRequest) (*notification.CreateNotificationResponse, error) {

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

func (c *Server) CloseConnections() {
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
