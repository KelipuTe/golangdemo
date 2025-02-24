package web

import "context"

// 服务关闭时需要执行的回调方法
type F8ShutdownCallback func(context.Context)

// 服务关闭时需要执行的回调方法
type I9ShutdownCallback interface {
	// 添加服务关闭时需要执行的回调方法
	F8AddShutdownCallback(...F8ShutdownCallback)
}
