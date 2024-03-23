package rpc

import (
	"context"
	"fmt"
	"os"

	"github.com/ljlimjk10/gym-avail-ms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCGetUpcomingTickets(ctx context.Context) ([]*queue.Ticket, error) {

	var tickets []*queue.Ticket
	grpcServerQueue := os.Getenv("GRPC_SERVER_QUEUE")
	if grpcServerQueue == "" {
		return nil, fmt.Errorf("missing .env url for queue grpc server")
	}

	qcc, err := grpc.Dial(grpcServerQueue, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection")
	}
	qClient := queue.NewQueueServiceClient(qcc)

	response, err := qClient.GetUpcomingTickets(ctx, &queue.GetUpcomingTicketsRequest{
		Quantity: 3,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get upcoming")
	}

	go func() {
		for _, ticket := range tickets {
			currentNumber := tickets[0].GetQueueNumber()
			GRPCSendNotification(ctx, &currentNumber, ticket, "Coming-Soon")
		}
	}()

	tickets = response.GetTickets()
	return tickets, nil
}
