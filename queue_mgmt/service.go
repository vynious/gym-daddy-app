package queue_mgmt

import (
	"context"
	"fmt"
	"github.com/vynious/gd-queue-ms/db"
)

type QueueService struct {
	*db.RedisDB
}

// Enqueue returns Ticket Number (current size) of the queue
func (q *QueueService) Enqueue(ctx context.Context, userId string) (int64, error) {
	number, err := q.RPush(ctx, q.QueueName, userId).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to enqueue %w", err)
	}
	return number, nil
}

// Dequeue retrieves the next userId (string) from the queue
func (q *QueueService) Dequeue(ctx context.Context, userId string) (string, error) {
	// when its someone's turn
	userId, err := q.LPop(ctx, q.QueueName).Result()
	if err != nil {
		return "", fmt.Errorf("failed to dequeue %w", err)
	}
	return userId, nil
}

func (q *QueueService) GetSize(ctx context.Context) (int64, error) {
	size, err := q.LLen(ctx, q.QueueName).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get size of queue %w", err)
	}
	return size, nil
}
