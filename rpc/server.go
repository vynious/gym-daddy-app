package rpc

import (
	"context"
	"github.com/vynious/gd-queue-ms/pb/proto_files/queue"
	"github.com/vynious/gd-queue-ms/queue_mgmt"
	"github.com/vynious/gd-queue-ms/util"
)

type Server struct {
	QueueMgmt *queue_mgmt.QueueService
	*queue.UnimplementedQueueServiceServer
}

func SpawnServer(qs *queue_mgmt.QueueService) *Server {
	return &Server{
		QueueMgmt: qs,
	}
}

func (s *Server) JoinQueue(ctx context.Context, req *queue.JoinQueueRequest) (*queue.JoinQueueResponse, error) {

	ticket, err := s.QueueMgmt.CreateTicket(ctx, req.UserID)

	if err != nil {
		return nil, err
	}
	if err = s.QueueMgmt.Enqueue(ctx, ticket); err != nil {
		return nil, err
	}
	if err = util.NotifyThroughTelegram(ticket.UserID); err != nil {
		return nil, err
	}

	return &queue.JoinQueueResponse{
		Ticket: ticket,
	}, nil
}

func (s *Server) RetrieveNext(ctx context.Context, req *queue.RetrieveNextRequest) (*queue.RetrieveNextResponse, error) {
	ticket, err := s.QueueMgmt.Dequeue(ctx)
	if err != nil {
		return nil, err
	}
	return &queue.RetrieveNextResponse{
		Ticket: ticket,
	}, nil
}

func (s *Server) GetUpcomingTickets(ctx context.Context, req *queue.GetUpcomingTicketsRequest) (*queue.GetUpcomingTicketsResponse, error) {
	tickets, err := s.QueueMgmt.RetrieveUpcoming(ctx, req.GetQuantity())
	if err != nil {
		return nil, err
	}
	return &queue.GetUpcomingTicketsResponse{
		Tickets: tickets,
	}, nil
}
