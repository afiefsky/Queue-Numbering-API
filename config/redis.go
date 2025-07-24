package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

const defaultRedisAddr = "localhost:6379"

func InitRedis(ctx context.Context) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     defaultRedisAddr,
		Password: "",
		DB:       0,
	})

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Redis ping error: %v", err)
	}

	log.Println("✅ Redis connected")
}
