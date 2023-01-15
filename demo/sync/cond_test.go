package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type S6MailBox struct {
	// 信箱状态，true=有信，false=没有信
	boxTag bool
	// 读写锁，保护 boxTag
	p7s6BoxRWLock *sync.RWMutex
	// 发信条件变量
	p7s6SendCond *sync.Cond
	// 收信条件变量
	p7s6ReceiveCond *sync.Cond
}

func f8NewS6MailBox() *S6MailBox {
	p7s6BoxRWLock := &sync.RWMutex{}
	// 发信时需要写锁，收信时只需要读锁
	return &S6MailBox{
		p7s6BoxRWLock:   p7s6BoxRWLock,
		p7s6SendCond:    sync.NewCond(p7s6BoxRWLock),
		p7s6ReceiveCond: sync.NewCond(p7s6BoxRWLock.RLocker()),
	}
}

func (p7this *S6MailBox) f8SendMail(i int) {
	time.Sleep(time.Millisecond * 500)
	p7this.p7s6BoxRWLock.Lock()
	for p7this.boxTag {
		p7this.p7s6SendCond.Wait()
	}
	fmt.Printf("send-%d:mailbox empty\n", i)
	p7this.boxTag = true
	fmt.Printf("send-%d:send mail\n", i)
	p7this.p7s6BoxRWLock.Unlock()
	// 唤醒一个收信的
	p7this.p7s6ReceiveCond.Signal()
}

func (p7this *S6MailBox) f8ReceiveMail(i int) {
	time.Sleep(time.Millisecond * 500)
	p7this.p7s6BoxRWLock.RLock()
	for !p7this.boxTag {
		p7this.p7s6ReceiveCond.Wait()
	}
	fmt.Printf("receive-%d:mailbox full\n", i)
	p7this.boxTag = false
	fmt.Printf("receive-%d:receive mail\n", i)
	p7this.p7s6BoxRWLock.RUnlock()
	// 唤醒一个发信的
	p7this.p7s6SendCond.Signal()
}

func TestSyncCond(p7s6t *testing.T) {
	testTime := 3
	p7s6MailBox := f8NewS6MailBox()
	for i := 1; i <= testTime; i++ {
		go p7s6MailBox.f8ReceiveMail(i)
	}
	for i := 1; i <= testTime; i++ {
		go p7s6MailBox.f8SendMail(i)
	}
	time.Sleep(5 * time.Second)
}
