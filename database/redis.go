package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	*redis.Client
}

func NewRedis(ctx context.Context) *Redis {
	client := connectRedis(ctx)

	return &Redis{client}
}

func connectRedis(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return client
}
