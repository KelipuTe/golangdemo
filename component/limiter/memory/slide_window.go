package memory

import (
	"context"
	"sync"
	"time"
)

// SlideWindowLimiter 滑动窗口限流器（内存实现）
type SlideWindowLimiter struct {
	windowSize time.Duration // 窗口总时长
	sliceSize  time.Duration // 每个时间片的时长
	maxReqNum  int           // 最大请求数
	rwLock     sync.RWMutex  // 互斥锁，锁 eachWindow
	eachWindow map[string]*slideWindowData
}

type slideWindowData struct {
	windowSize    time.Duration   // 窗口总时长
	sliceSize     time.Duration   // 每个时间片的时长
	maxReqNum     int             // 最大请求数
	lock          sync.Mutex      // 互斥锁，锁 slideWindowData
	eachSlice     []timeSliceData // 每个时间片的数据
	lastCleanTime int64           //最后清理时间
}

type timeSliceData struct {
	endTime int64 // 时间片结束时间
	count   int   // 请求数
}

func NewSlideWindowLimiter(w, s time.Duration, num int) *SlideWindowLimiter {
	return &SlideWindowLimiter{
		windowSize: w,
		sliceSize:  s,
		maxReqNum:  num,
		eachWindow: map[string]*slideWindowData{},
	}
}

func (t *SlideWindowLimiter) IsLimit(ctx context.Context, key string) (bool, error) {
	now := time.Now().UnixNano()

	// 运行起来之后，读肯定远多于写，所以用读写锁
	t.rwLock.RLock()
	window, exist := t.eachWindow[key]
	t.rwLock.RUnlock()

	// 当限流对象没有对应的限流器时，创建一个
	if !exist {
		window = &slideWindowData{
			windowSize:    t.windowSize,
			sliceSize:     t.sliceSize,
			maxReqNum:     t.maxReqNum,
			lock:          sync.Mutex{},
			eachSlice:     []timeSliceData{},
			lastCleanTime: now,
		}
		t.rwLock.Lock()
		t.eachWindow[key] = window
		t.rwLock.Unlock()
	}

	return window.isLimit(ctx, now)
}

func (t *slideWindowData) isLimit(ctx context.Context, now int64) (bool, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	//计算窗口开始时间
	windowStart := now - t.windowSize.Nanoseconds()

	//清理过期的时间片，每经过一个时间片，就要清理一次
	if now-t.lastCleanTime > t.sliceSize.Nanoseconds() {
		t.cleanTimeSliceData(windowStart)
		t.lastCleanTime = now
	}

	//计算窗口总请求数
	total := 0
	for _, s := range t.eachSlice {
		total += s.count
	}

	//如果触发限流就直接返回
	if total >= t.maxReqNum {
		return true, nil
	}

	//计算当前时间片的结束时间
	nowEndTime := ((now / t.sliceSize.Nanoseconds()) + 1) * t.sliceSize.Nanoseconds()

	//记录请求次数
	if len(t.eachSlice) > 0 && t.eachSlice[len(t.eachSlice)-1].endTime == nowEndTime {
		t.eachSlice[len(t.eachSlice)-1].count++
	} else {
		t.eachSlice = append(t.eachSlice, timeSliceData{
			endTime: nowEndTime,
			count:   1,
		})
	}

	return false, nil
}

func (t *slideWindowData) cleanTimeSliceData(windowStart int64) {
	//找到第一个时间片结束时间大于 windowStart 的时间片，前面的都是过期的
	firstIndex := 0
	for ; firstIndex < len(t.eachSlice); firstIndex++ {
		if t.eachSlice[firstIndex].endTime > windowStart {
			break
		}
	}
	t.eachSlice = t.eachSlice[firstIndex:]
}
