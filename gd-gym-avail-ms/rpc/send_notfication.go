package rpc

import (
	"context"
	"log"
	"os"

	"github.com/ljlimjk10/gym-avail-ms/pb/proto_files/notification"
	"github.com/ljlimjk10/gym-avail-ms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCSendNotification(ctx context.Context, currentNumber *int64, ticket *queue.Ticket, nType string) {

	grpcServerNotification := os.Getenv("GRPC_SERVER_NOTIFICATION")
	if grpcServerNotification == "" {
		log.Println("missing .env url for notification grpc server")
	}

	ncc, err := grpc.Dial(grpcServerNotification, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("failed to create grpc connection")
	}

	nClient := notification.NewNotificationServiceClient(ncc)

	_, err = nClient.CreateNotification(ctx, &notification.CreateNotificationRequest{
		CurrentQueueNumber: currentNumber,
		UserTicket:         ticket,
		NotificationType:   nType,
	})

	if err != nil {
		log.Println(err.Error())
		log.Println("failed to send notification")
	}

}
