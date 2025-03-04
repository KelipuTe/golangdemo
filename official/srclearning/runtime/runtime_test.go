package runtime

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

// runtime 包可以获得 go 运行时的一些信息
// 比如：go 的版本；机器的cpu数量；goroutine的数量；程序的调用栈信息；内存使用情况；等；

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
			// 会输出 runtime_test.go 32 true，
			// 就是上面这行 runtime.Caller(0) 在哪个文件的多少行
			// true 这表示该函数确实存在于源码中

			log.Println("runtime.Caller(1):")
			fmt.Println(runtime.Caller(1))
			// 会输出 /srclearning/runtime/panic.go 884 true
			// 这个是 go 标准库 runtime 包的一部分

			log.Println("runtime.Caller(2):")
			fmt.Println(runtime.Caller(2))
			// 会输出 runtime_test.go 55 true，
			// 就是上面的 recover 捕获的 panic 在哪个文件的多少行
			// true 这表示该函数确实存在于源码中

			log.Println("runtime.Caller(3):")
			fmt.Println(runtime.Caller(3))
			// 会输出 /srclearning/testing/testing.go 1576 true
			// 这个是 go 标准库 testing 包的一部分
		}
	}()

	panic("TestStackAndCaller")
}

func TestMemory(t *testing.T) {
	log.Println("TestMemory:")
	printMemStats()

	for i := 1; i <= 5; i++ {
		// 申请一个 10 MB 的切片
		_ = make([]byte, 10*1024*1024)
		log.Println("for:", i)
		printMemStats()
	}

	runtime.GC() // 显式运行垃圾回收
	log.Println("runtime.GC:")
	printMemStats()
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %d MB", bToMb(int(m.Alloc)))
	fmt.Printf("\tSys = %d MB", bToMb(int(m.Sys)))
	fmt.Printf("\tNumGC = %d\n", m.NumGC)
}

// 字节转兆字节
func bToMb(num int) int {
	return num / 1024 / 1024
}
