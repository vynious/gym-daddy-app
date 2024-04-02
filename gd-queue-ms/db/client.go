package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

type RedisDB struct {
	*redis.Client
	QueueName string
}

func SpawnRedisDB(queueName string) (*RedisDB, error) {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		return nil, fmt.Errorf("missing redis url in .gitignore")
	}
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url")
	}
	r := redis.NewClient(opts)
	_, err = r.Set(context.Background(), "ticket_count", 0, 0).Result()
	if err != nil {
		return nil, fmt.Errorf("fail to init counter %w", err)
	}
	return &RedisDB{
		r,
		queueName,
	}, nil
}
