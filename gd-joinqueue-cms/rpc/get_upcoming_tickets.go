package rpc

import (
	"context"
	"fmt"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GRPCGetUpcomingTickets(ctx context.Context, sendNotification bool) ([]*queue.Ticket, error) {

	var tickets []*queue.Ticket
	grpcServerQueue := os.Getenv("GRPC_SERVER_QUEUE")
	if grpcServerQueue == "" {
		return nil, fmt.Errorf("missing .gitignore url for queue grpc server")
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

	if sendNotification {
		go func() {
			for _, ticket := range tickets {
				currentNumber := ticket.GetQueueNumber()
				GRPCSendNotification(ctx, &currentNumber, ticket, "Coming-Soon")
			}
		}()
	}

	tickets = response.GetTickets()
	return tickets, nil
}
