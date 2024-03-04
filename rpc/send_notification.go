package rpc

import (
	"context"
	"fmt"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/notification"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RPCSendNotification(ctx context.Context, ticket *queue.Ticket) {
	ncc, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("failed to create grpc connection")
	}
	nClient := notification.NewNotificationServiceClient(ncc)
	_, err = nClient.CreateNotification(ctx, &notification.CreateNotificationRequest{
		TelegramHandle:   "",
		NotificationType: "Join-Queue",
		Content:          fmt.Sprintf("You have joined the queue! Your ticket number is %v", ticket.QueueNumber),
	})
	if err != nil {
		log.Println("failed to send notification")
	}
}
