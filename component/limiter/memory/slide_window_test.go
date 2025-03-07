package memory

import (
	"context"
	"testing"
	"time"
)

func TestSlideWindowLimiter(t *testing.T) {
	limiter := NewSlideWindowLimiter(3*time.Second, time.Second, 3)

	go func() {
		for {
			time.Sleep(time.Second)
			ctx := context.Background()
			isLimit, err := limiter.IsLimit(ctx, "test")
			t.Log("test", isLimit, err)
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			time.Sleep(time.Second)
			ctx := context.Background()
			isLimit, err := limiter.IsLimit(ctx, "test")
			t.Log("test", isLimit, err)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			ctx := context.Background()
			isLimit, err := limiter.IsLimit(ctx, "test2")
			t.Log("test2", isLimit, err)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
