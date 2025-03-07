package redis

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed slide_window.lua
var luaScript string

// SlideWindowLimiter 滑动窗口限流器（redis 实现）
type SlideWindowLimiter struct {
	redis      redis.Cmdable
	windowSize time.Duration // 窗口总时长
	maxReqNum  int           // 最大请求数
}

func NewSlideWindowLimiter(cmd redis.Cmdable, w time.Duration, num int) *SlideWindowLimiter {
	return &SlideWindowLimiter{
		redis:      cmd,
		windowSize: w,
		maxReqNum:  num,
	}
}

func (t *SlideWindowLimiter) IsLimit(ctx context.Context, key string) (bool, error) {
	return t.redis.Eval(
		ctx, luaScript, []string{key},
		t.windowSize.Milliseconds(), t.maxReqNum, time.Now().UnixMilli(),
	).Bool()
}
