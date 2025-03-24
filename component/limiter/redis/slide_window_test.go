package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestSlideWindowLimiter(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	limiter := NewSlideWindowLimiter(rdb, 3*time.Second, 3)

	go func() {
		for {
			time.Sleep(time.Second)
			ctx := context.Background()
			isLimit, err := limiter.IsLimited(ctx, "test")
			t.Log("test", isLimit, err)
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			time.Sleep(time.Second)
			ctx := context.Background()
			isLimit, err := limiter.IsLimited(ctx, "test")
			t.Log("test", isLimit, err)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			ctx := context.Background()
			isLimit, err := limiter.IsLimited(ctx, "test2")
			t.Log("test2", isLimit, err)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
