package util

import (
	"context"
	"fmt"
	"github.com/vynious/gd-queue-ms/pb/proto_files/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NotifyThroughTelegram(userID string) error {

	var n notification.Notification

	// call user-management-ms to get telegram handle to create the notification

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
