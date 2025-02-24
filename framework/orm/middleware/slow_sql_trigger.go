package middleware

import (
	"context"
	"demo-golang/orm"
	"time"
)

// SlowLogTriggerMiddlewareBuild 触发慢 SQL 用的
func SlowLogTriggerMiddlewareBuild() orm.F8Middleware {
	return func(next orm.F8MiddlewareHandle) orm.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *orm.S6QueryContext) *orm.S6QueryResult {
			time.Sleep(500 * time.Millisecond)
			return next(ctx, p7s6Context)
		}
	}
}
