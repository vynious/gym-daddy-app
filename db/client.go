package db

import (
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
		return nil, fmt.Errorf("missing redis url in .env")
	}
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url")
	}
	r := redis.NewClient(opts)

	return &RedisDB{
		r,
		queueName,
	}, nil
}
