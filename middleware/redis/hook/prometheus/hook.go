package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"log"
	"net"
	"time"
)

// Hook 给GoRedis添加Prometheus插件
// 监控，命令执行时间
type Hook struct {
	vector *prometheus.SummaryVec
}

func NewHook(vector *prometheus.SummaryVec) *Hook {
	return &Hook{
		vector: vector,
	}
}

// DialHook 创建网络连接时调用的hook
func (t *Hook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

// ProcessHook 执行命令时调用的hook
func (t *Hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		log.Println("ProcessHook", cmd.Name(), cmd.Args())

		var err error

		startTime := time.Now()
		defer func() {
			//本地速度太快，毫秒有可能是0，换微秒方便观察
			timeSince := time.Since(startTime).Microseconds()

			pattern, patternOk := ctx.Value("pattern").(string)
			if !patternOk {
				pattern = "pattern"
			}

			log.Println("ProcessHook defer", cmd.Name(), pattern, timeSince)

			t.vector.WithLabelValues(cmd.Name(), pattern).Observe(float64(timeSince))
		}()

		err = next(ctx, cmd)

		return err
	}
}

// ProcessPipelineHook 执行管道命令时调用的hook
func (t *Hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}
