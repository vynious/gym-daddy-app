package rpc

import (
	"context"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/notification"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func GRPCSendNotification(ctx context.Context, currentNumber *int64, ticket *queue.Ticket, nType string) {

	grpcServerNotification := os.Getenv("GRPC_SERVER_NOTIFICATION")
	if grpcServerNotification == "" {
		log.Println("missing .env url for notification grpc server")
		return
	}

	// Establish a connection to the gRPC server
	ncc, err := grpc.Dial(grpcServerNotification, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("failed to create grpc connection: %v", err)
		return
	}
	defer ncc.Close()

	nClient := notification.NewNotificationServiceClient(ncc)

	// Send a CreateNotification request
	_, err = nClient.CreateNotification(ctx, &notification.CreateNotificationRequest{
		UserTicket:         ticket,
		CurrentQueueNumber: currentNumber,
		NotificationType:   nType,
	})

	if err != nil {
		log.Printf("failed to send notification: %v", err)
		return
	}

	log.Println("Notification sent successfully")
}