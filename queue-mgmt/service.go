package queue_mgmt

import (
	"context"
	"fmt"
	"github.com/vynious/gd-queue-ms/db"
)

type Queue struct {
	*db.RedisDB
}

func (q *Queue) Enqueue(ctx context.Context, userId string) error {
	if err := q.RPush(ctx, q.QueueName, userId).Err(); err != nil {
		return fmt.Errorf("failed to enqueue %w", err)
	}
	return nil
}

// Dequeue Retrieve the next userId (string) from the queue
func (q *Queue) Dequeue(ctx context.Context, userId string) (string, error) {
	userId, err := q.LPop(ctx, q.QueueName).Result()
	if err != nil {
		return "", fmt.Errorf("failed to dequeue %w", err)
	}
	return userId, nil
}

func (q *Queue) GetSize(ctx context.Context) (int64, error) {
	size, err := q.LLen(ctx, q.QueueName).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get size of queue %w", err)
	}
	return size, nil
}
