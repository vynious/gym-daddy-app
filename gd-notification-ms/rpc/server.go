package rpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vynious/gym-daddy/gd-notification-ms/db"
	"github.com/vynious/gym-daddy/gd-notification-ms/kafka"
	"github.com/vynious/gym-daddy/gd-notification-ms/pb/proto_files/notification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	var notificationProto notification.Notification

	if len(notificationProto.Id) > 0 {

	} else {
		id, err := uuid.NewUUID()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate uuid for notification: %v", err)
		}
		notificationProto.Id = id.String()
	}

	notificationProto.NotificationType = req.GetNotificationType()
	notificationProto.CreatedAt = timestamppb.Now()

	userTicket := req.GetUserTicket()
	if userTicket == nil {
		return nil, fmt.Errorf("missing ticket in request")
	}
	// todo: make call to get telegram_handle
	notificationProto.TelegramHandle = "shawntyw" // default

	switch req.GetNotificationType() {
	case "Join-Queue":
		notificationProto.Content = fmt.Sprintf("You have joined the queue! Your ticket number is %v", userTicket.QueueNumber)
	case "Coming-Soon":
		notificationProto.Content = fmt.Sprintf("It's almost your turn soon! Prepare to come down ~")
	}

	errCh := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.Repository.SaveNotification(&notificationProto); err != nil {
			errCh <- status.Errorf(codes.Internal, "failed to save notitfication to db: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.Notifier.SendNotification(ctx, &notificationProto); err != nil {
			errCh <- status.Errorf(codes.Internal, "failed to send notitfication to kafka queue: %v", err)
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		return nil, err
	}

	return &notification.CreateNotificationResponse{
		Notification: &notificationProto,
	}, nil
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
