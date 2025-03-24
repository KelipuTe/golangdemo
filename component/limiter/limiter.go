package limiter

import "context"

type Limiter interface {
	// IsLimited true=触发限流；false=未触发限流
	IsLimited(ctx context.Context, key string) (bool, error)
}
