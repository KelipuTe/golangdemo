package signal

import (
	"fmt"
	"os"
	"os/signal"
)

// WaitForSIGINT 按下 ctrl+c 时，触发的就是 SIGINT 信号
func WaitForSIGINT() {
	fmt.Println("use ctrl+c")
	//管道不设置缓冲区有可能会错过信号
	waitChannel := make(chan os.Signal, 1)
	//os.Interrupt 就是 SIGINT
	signal.Notify(waitChannel, os.Interrupt)
	//如果没有收到信号，就会阻塞在这里
	s := <-waitChannel
	fmt.Println("get signal:", s)
}
