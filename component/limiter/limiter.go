package limiter

import "context"

type Limiter interface {
	// IsLimit true=触发限流；false=未触发限流
	IsLimit(ctx context.Context, key string) (bool, error)
}
