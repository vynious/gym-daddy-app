package rpc

import (
	"context"
	"fmt"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GRPCJoinQueue(ctx context.Context, userId string) (*queue.Ticket, error) {

	var ticket *queue.Ticket

	grpcServerQueue := os.Getenv("GRPC_SERVER_QUEUE")
	if grpcServerQueue == "" {
		return nil, fmt.Errorf("missing .env url for queue grpc server")
	}
	qcc, err := grpc.Dial(grpcServerQueue, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection")
	}
	qClient := queue.NewQueueServiceClient(qcc)
	response, err := qClient.JoinQueue(ctx, &queue.JoinQueueRequest{
		UserId: userId,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to join queue")
	}
	ticket = response.GetTicket()

	go GRPCSendNotification(ctx, nil, ticket, "Join-Queue")

	return ticket, nil
}
