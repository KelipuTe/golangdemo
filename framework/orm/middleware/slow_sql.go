package middleware

import (
	"context"
	"demo-golang/orm"
	"log"
	"time"
)

// SlowLogMiddlewareBuild 计算查询执行时间，用于捕获慢 SQL
func SlowLogMiddlewareBuild() orm.F8Middleware {
	return func(next orm.F8MiddlewareHandle) orm.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *orm.S6QueryContext) *orm.S6QueryResult {
			timeStart := time.Now()
			t4 := next(ctx, p7s6Context)
			timeEnd := time.Now()
			timeCost := timeEnd.Sub(timeStart).Milliseconds()
			log.Printf("time pass %d ms\r\n", timeCost)
			if 200 < timeCost {
				log.Printf("slow sql, time pass %d ms\r\n", timeCost)
			}
			return t4
		}
	}
}
