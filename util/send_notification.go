package util

import (
	"context"
	"fmt"
	"github.com/vynious/gd-queue-ms/pb/proto_files/notification"
	"github.com/vynious/gd-queue-ms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NotifyThroughTelegram(ticket *queue.Ticket) error {

	var n notification.Notification

	n.Type = "queue"
	n.CreatedAt = timestamppb.Now()
	n.Content = fmt.Sprintf("Congrats! Your queue number is %v", ticket.QueueNumber)

	// todo: call user-management-ms to get telegram handle to create the notification

	cc, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return fmt.Errorf("failed to make grpc dial to server")
	}
	c := notification.NewNotificationServiceClient(cc)

	defer cc.Close()

	// GRPC dial to server
	_, err = c.CreateNotification(context.Background(), &notification.CreateNotificationRequest{
		Notification: &n,
	})

	if err != nil {
		return fmt.Errorf("failed to send notification to notification ms: %w", err)
	}

	return nil
}
