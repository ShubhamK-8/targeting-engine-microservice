package redis

import (
	"context"
	"fmt"
	"time"

	coreUtils "targeting-engine/coreUtils"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

// Initializes and returns a new Redis client.
func NewRedisClient() (*RedisClient, error) {
	// Replace with your Redis address
	redisClient := redis.NewClient(&redis.Options{
		Addr:     coreUtils.RedisHost, // Default Redis port
		Password: "",                  // No password set
		DB:       0,                   // Default DB
	})

	// Test the Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Errorf("Could not connect to Redis: %v", err)
		return nil, err
	}
	fmt.Println("Successfully connected to Redis!")
	return &RedisClient{Client: redisClient}, nil
}
