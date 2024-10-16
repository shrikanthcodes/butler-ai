package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
)

type ConnPool struct {
	*redis.Client
}

// New initializes the Redis client.
func New(cfg config.Redis, log logger.Logger) (*ConnPool, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // No password set
		DB:       cfg.DB,       // Use default DB
		PoolSize: cfg.PoolMax,  // Set max pool size
	})

	// Test the connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Info("Connected to Redis at", cfg.Address)
	return &ConnPool{rdb}, nil
}
