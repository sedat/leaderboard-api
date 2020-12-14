package tests

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx context.Context
var redisClient *redis.Client

func init() {
	ctx = context.Background()
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func TestRedisConnection(t *testing.T) {
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		t.Errorf("Redis connection error")
	}
}
