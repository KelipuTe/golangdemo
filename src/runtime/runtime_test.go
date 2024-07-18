package runtime

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	log.Println("go的版本号=", runtime.Version())
	log.Println("cpu的数量=", runtime.NumCPU())
	log.Println("当前goroutine的数量=", runtime.NumGoroutine())
}

func TestStackAndCaller(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// 获取调用栈信息
			stackTrace := make([]byte, 4096)
			stackTraceLen := runtime.Stack(stackTrace, false)

			log.Println("panic:", err)

			log.Println("runtime.Stack:")
			fmt.Println(string(stackTrace[:stackTraceLen]))

			log.Println("runtime.Caller(0):")
			fmt.Println(runtime.Caller(0))
			// 会输出 runtime_test.go 29 true
			// 就是调用 runtime.Caller(0) 的位置，true 这表示该函数确实存在于源码中

			log.Println("runtime.Caller(1):")
			fmt.Println(runtime.Caller(1))
			// 会输出 /src/runtime/panic.go 884 true
			// 这个是 go 标准库 runtime 包的一部分

			log.Println("runtime.Caller(2):")
			fmt.Println(runtime.Caller(2))
			// 会输出 runtime_test.go 50 true
			// 就是触发 panic 的位置，true 这表示该函数确实存在于源码中

			log.Println("runtime.Caller(3):")
			fmt.Println(runtime.Caller(3))
			// 会输出 /src/testing/testing.go 1576 true
			// 这个是 go 标准库 testing 包的一部分
		}
	}()

	panic("TestStackAndCaller")
}
