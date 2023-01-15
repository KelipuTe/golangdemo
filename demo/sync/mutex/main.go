package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

//互斥量或者互斥锁（syncMutextual exclusion，简称 syncMutex）

//不要重复锁定，不要忘记解锁，不要对尚未锁定或者已解锁的锁解锁。

//不要在多个函数之间直接传递锁，互斥锁一个结构体类型，属于值类型中的一种，
//把它传给一个函数、将它从函数中返回、把它赋给其他变量、让它进入某个通道，都会导致副本的产生。
//原值和它的副本，以及多个副本之间都是完全独立的，它们都是不同的互斥锁。

func main() {
	var (
		gortNum     int           = 5    //并发写的goroutine数量
		writeTime   int           = 5    //每个goroutine写入几次
		numTime     int           = 10   //每次写入几个数字
		chanSign    chan struct{}        //信号通道
		writeBuffer bytes.Buffer         //缓冲区
		writeLock   sync.Mutex           //互斥锁
		useLock     bool          = true //是否启用互斥锁

		funcGortWrite func(gortId int, writer io.Writer) //并发写的函数
	)

	chanSign = make(chan struct{}, gortNum)

	funcGortWrite = func(gortId int, writer io.Writer) {
		defer func() { chanSign <- struct{}{} }()

		for i := 1; i <= writeTime; i++ {
			//提示文本和重复文本
			header := fmt.Sprintf("\ngort id:%d,write time:%d:", gortId, i)
			data := fmt.Sprintf(" %d", gortId)

			if useLock {
				writeLock.Lock()
			}

			_, err := writer.Write([]byte(header))
			if err != nil {
				log.Printf("error: %s [%d]", err, gortId)
			}

			for j := 0; j < numTime; j++ {
				_, err := writer.Write([]byte(data))
				if err != nil {
					log.Printf("error: %s [%d]", err, gortId)
				}
			}

			if useLock {
				writeLock.Unlock()
			}
		}
	}

	//执行并发写
	for i := 1; i <= gortNum; i++ {
		go funcGortWrite(i, &writeBuffer)
	}
	//等待全部goroutine结束
	for i := 0; i < gortNum; i++ {
		<-chanSign
	}
	//读取和输出缓冲区中的数据
	data, err := ioutil.ReadAll(&writeBuffer)
	if err != nil {
		fmt.Printf("error:%s\n", err)
	}
	fmt.Printf("writeBuffer:\n%s\n", data)
}
