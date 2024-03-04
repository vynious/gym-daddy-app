package rpc

import (
	"context"
	"github.com/vynious/gd-joinqueue-cms/pb/proto_files/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RPCJoinQueue(ctx context.Context, userId string) error {

	var ticket *queue.Ticket

	qcc, err := grpc.Dial("localhost:3002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("failed to create grpc connection")
	}
	qClient := queue.NewQueueServiceClient(qcc)
	response, err := qClient.JoinQueue(ctx, &queue.JoinQueueRequest{
		UserId: userId,
	})
	if err != nil {
		log.Println("failed to join queue")
	}
	ticket = response.GetTicket()

	go RPCSendNotification(ctx, ticket)

	return nil
}
