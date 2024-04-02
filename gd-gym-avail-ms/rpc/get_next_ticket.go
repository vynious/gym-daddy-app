package rpc

import (
	"context"
	"fmt"
	"os"

	"github.com/ljlimjk10/gym-avail-ms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCGetNextInQueue(ctx context.Context) (*queue.Ticket, error) {

	var ticket *queue.Ticket

	grpcServerQueue := os.Getenv("GRPC_SERVER_QUEUE")
	if grpcServerQueue == "" {
		return nil, fmt.Errorf("missing .gitignore url for queue grpc server")
	}

	qcc, err := grpc.Dial(grpcServerQueue, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to make grpc conn")
	}

	qClient := queue.NewQueueServiceClient(qcc)

	response, err := qClient.RetrieveNext(ctx, &queue.RetrieveNextRequest{})

	ticket = response.GetTicket()

	if err != nil {
		return nil, fmt.Errorf("failed to get next ticket")
	}

	return ticket, nil
}
