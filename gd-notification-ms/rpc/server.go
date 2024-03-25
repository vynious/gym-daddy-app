package rpc

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/vynious/gym-daddy/gd-notification-ms/db"
	"github.com/vynious/gym-daddy/gd-notification-ms/kafka"
	"github.com/vynious/gym-daddy/gd-notification-ms/pb/proto_files/notification"
	"github.com/vynious/gym-daddy/gd-notification-ms/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	fmt.Println("receiving create notification request")

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

	telegramHandle, err := util.GetTelegramHandleFromUserMS(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get telegram handle: %v", err)
	}

	notificationProto.TelegramHandle = telegramHandle

	fmt.Printf("partial notification proto created %+v", &notificationProto)

	switch req.GetNotificationType() {
	case "Join-Queue":
		userTicket := req.GetUserTicket()
		if userTicket == nil {
			return nil, fmt.Errorf("missing ticket in request")
		}
		notificationProto.Content = fmt.Sprintf("You have joined the queue! Your ticket number is %v", userTicket.QueueNumber)
	case "Coming-Soon":
		notificationProto.Content = fmt.Sprintln("It's almost your turn soon! Prepare to come down ~")
	case "Booking-Confirmation":
		notificationProto.Content = fmt.Sprintln("We would like to notify you that you've made a new class booking, please refer to your account find out more.")
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
