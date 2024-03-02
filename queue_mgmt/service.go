package queue_mgmt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vynious/gd-queue-ms/db"
	"github.com/vynious/gd-queue-ms/pb/proto_files/queue"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type QueueService struct {
	*db.RedisDB
}

func SpawnQueueService(db *db.RedisDB) *QueueService {
	return &QueueService{
		db,
	}
}

func (q *QueueService) Enqueue(ctx context.Context, ticket *queue.Ticket) error {
	serializedTicket, err := json.Marshal(ticket)
	if err != nil {
		return fmt.Errorf("failed to serialize ticket: %w", err)
	}

	_, err = q.RPush(ctx, q.QueueName, serializedTicket).Result()
	if err != nil {
		return fmt.Errorf("failed to enqueue: %w", err)
	}

	return nil
}

func (q *QueueService) Dequeue(ctx context.Context) (*queue.Ticket, error) {
	serializedTicket, err := q.LPop(ctx, q.QueueName).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to dequeue: %w", err)
	}
	// Assuming the ticket data is stored as a JSON string in the queue,
	// we need to deserialize it back into a queue.Ticket object.
	var ticket queue.Ticket
	err = json.Unmarshal([]byte(serializedTicket), &ticket)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize ticket: %w", err)
	}
	return &ticket, nil
}

func (q *QueueService) CreateTicket(ctx context.Context, userId string) (*queue.Ticket, error) {
	// Start a Redis transaction.
	pipeline := q.TxPipeline()

	// Increment the ticket_count atomically.
	incr := pipeline.Incr(ctx, "ticket_count")

	// Execute the transaction.
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	// Create a new ticket with the incremented ticket_count as the QueueNumber.
	ticket := &queue.Ticket{
		UserID:      userId,
		QueueNumber: incr.Val(), // Use the result of the INCR operation as the queue number.
		CreatedAt:   timestamppb.Now(),
	}
	return ticket, nil
}

func (q *QueueService) RetrieveUpcoming(ctx context.Context, quantity int64) ([]*queue.Ticket, error) {
	var ticketList []*queue.Ticket
	tickets, err := q.LRange(ctx, q.QueueName, 0, quantity-1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get upcoming tickets")
	}
	for _, t := range tickets {
		var ticket queue.Ticket
		err = json.Unmarshal([]byte(t), &ticket)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize ticket: %w", err)
		}
		ticketList = append(ticketList, &ticket)
	}
	return ticketList, nil
}
