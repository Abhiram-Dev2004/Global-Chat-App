package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func SetupRedisClient() (*redis.Client, *redis.PubSub) {
	ctx := context.Background()
	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASS")
	opt := &redis.Options{
		Addr:      redisHost,
		Username:  "default",
		Password:  redisPass,
		DB:        0,
		TLSConfig: &tls.Config{}, // Required for `rediss://`
	}

	client := redis.NewClient(opt)

	// Test connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("!!!!Connected to Redis:", pong)

	// Pub/Sub setup

	sub := client.Subscribe(ctx, "my-channel") // Subscriber listens to same channel

	return client, sub
}
