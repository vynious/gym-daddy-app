package rpc

import (
	"context"
	"github.com/vynious/gd-queue-ms/pb/proto_files/queue"
	"github.com/vynious/gd-queue-ms/queue_mgmt"
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
}
