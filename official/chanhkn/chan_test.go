package chanhkn

import (
	"sync"
	"testing"
	"time"
)

// 没缓存的管道，读写必须同时进行，否则阻塞
func TestSyncChan(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	val := <-ch
	t.Log(val)

	close(ch)
}

// 写一个已关闭的管道会panic
func TestWriteCloseChanPanic(t *testing.T) {
	ch := make(chan int)
	close(ch)
	ch <- 1
}

// 关闭一个已经关闭的管道会panic
func TestCloseClosedChanPanic(t *testing.T) {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// 有缓存的管道，读写可以分开进行
// 如果管道为空，读会阻塞；如果缓存满了，写会阻塞；
// 管道永久阻塞是导致 goroutine 泄露的常见原因
func TestCacheChan(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1
	val := <-ch
	t.Log(val)

	close(ch)
}

// 没缓存的管道关闭时，会读到0值和false
func TestReadCloseChan(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		close(ch)
	}()

	val, ok := <-ch
	t.Log(val, ok)
}

// 有缓存的管道关闭时，先读到数据和true，后读到0值和false
func TestReadCloseCacheChan(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1
	val, ok := <-ch
	t.Log(val, ok)

	close(ch)

	val, ok = <-ch
	t.Log(val, ok)
}

// 用管道传递信号
func TestSendSignal(t *testing.T) {
	//声明的时候一般用空结构体
	ch := make(chan struct{}, 1)

	ch <- struct{}{}
	val, ok := <-ch
	t.Log(val, ok)

	close(ch)
}

type SafeCloseChan struct {
	ch        chan int  //一定是私有的
	closeOnce sync.Once //确保只执行一次
}

func (s *SafeCloseChan) Close() error {
	s.closeOnce.Do(func() {
		close(s.ch)
	})
	return nil
}

// 可以安全关闭的管道
func TestSafeCloseChan(t *testing.T) {
	ch := &SafeCloseChan{
		ch: make(chan int, 1),
	}
	_ = ch.Close()
	_ = ch.Close()
}

// 遍历管道，用 for range 就可以
func TestLoopChan(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 200)
		}
		close(ch)
	}()

	for val := range ch {
		t.Log(val)
	}

	t.Log("管道被关闭")
}

// 管道经常和select一起用
// 哪个case后面的语句没有阻塞，就走哪个case
// 如果多个case都没有阻塞，则随机选一个case执行
// 没有被选的case，case后面的代码不会执行
func TestUseWithSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	select {
	case val1 := <-ch1:
		t.Log("ch1", val1)
		val2 := <-ch2
		t.Log("ch2", val2)
	case val2 := <-ch2:
		t.Log("ch2", val2)
		val1 := <-ch1
		t.Log("ch1", val1)
	}
}

// 管道可以被声明为只读（常用），也可以声明为只写（不常用）
func TestReadOnlyChan(t *testing.T) {
	ch := returnReadOnlyChan()
	for val := range ch {
		t.Log(val)
	}
}

func returnReadOnlyChan() <-chan int {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	return ch
}
