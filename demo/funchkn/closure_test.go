package funchkn

//闭包的一些用法样例
//1 在闭包的内部可以修改从外部引用的变量
//2 闭包存在记忆效应，被捕获到闭包中的变量，会跟随闭包生命周期一直存在
//2.1 可以用来维护函数内部的状态，而无需显式地通过参数传递状态信息，比如计数器、累加器
//2.2 闭包可以用来实现延迟计算和缓存，使得某些计算只在需要时才执行，并缓存结果以供后续使用
//2.3 可以用来构造事件处理机制，在事件的触发和运行可以分开，事件运行时拿到的依然是事件触发时的状态

import (
	"log"
	"testing"
)

// 修改从外部引用的变量
func TestOutside(t *testing.T) {
	outside := 0
	closure := func() {
		outside++
		log.Println(outside)
	}

	closure()
	closure()
}

// 维护函数内部的状态
func TestStatus(t *testing.T) {
	closure := closureStatus()

	closure()
	closure()
}

func closureStatus() func() int {
	status := 0
	return func() int {
		status++
		log.Println(status)
		return status
	}
}

// 延迟计算和缓存
func TestCache(t *testing.T) {
	closure := closureCache(func(i int) int { return i * i })

	closure(4)
	closure(4)
}

func closureCache(delay func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if v, ok := cache[x]; ok {
			log.Println("cache", v)
			return v
		}
		result := delay(x)
		cache[x] = result
		log.Println("delay", result)
		return result
	}
}

// 构造事件处理机制
func TestEvent(t *testing.T) {
	data := "a"
	closure01 := closureEvent(data) // 生成闭包方法时，data 是 a
	data = "b"                      // 这里修改不会影响 closure01 里的 data
	closure02 := closureEvent(data) // 生成闭包方法时，data 是 b
	data = "c"                      // 这里修改不会影响 closure02 里的 data

	closure01()
	closure02()
}

func closureEvent(data string) func() {
	return func() {
		log.Println(data)
	}
}
