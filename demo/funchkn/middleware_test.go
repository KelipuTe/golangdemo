package funchkn

import (
	"log"
	"testing"
)

// 中间件样例

func TestMiddleware(t *testing.T) {
	middlewareList := []middlewareExample{
		logBeforeMiddleware(),
		logAfterMiddleware(),
	}

	chain := coreHandle
	//倒过来组装，先组装的在里层，里层的后执行
	for i := len(middlewareList) - 1; i >= 0; i-- {
		chain = middlewareList[i](chain)
	}
	ctx := &contextExample{}
	chain(ctx)
}

// 上下文
type contextExample struct{}

// 最内层的核心处理方法
func coreHandle(ctx *contextExample) {
	log.Println("coreHandle")
}

// 中间件结构
type middlewareHandle func(*contextExample)
type middlewareExample func(middlewareHandle) middlewareHandle

func logBeforeMiddleware() middlewareExample {
	return func(next middlewareHandle) middlewareHandle {
		return func(ctx *contextExample) {
			log.Println("logBefore")
			next(ctx)
		}
	}
}

func logAfterMiddleware() middlewareExample {
	return func(next middlewareHandle) middlewareHandle {
		return func(ctx *contextExample) {
			next(ctx)
			log.Println("logAfter")
		}
	}
}
