package runtime

import (
	"log"
	"testing"
	"time"
)

func TestFuncWithRecover(t *testing.T) {
	go func() {
		for {
			res := funcWithRecover()
			log.Println("res=", res)
			time.Sleep(time.Second)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}

func funcWithRecover() bool {
	//这个结构可以防止函数内部的panic导致函数外部挂掉
	defer func() { recover() }()

	unix := time.Now().Unix()
	if unix%10 == 0 {
		//如果走到这里触发panic了，外部拿到的会是bool的零值，也就是false
		panic("panic")
	}

	return true
}
