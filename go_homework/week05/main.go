package main

import (
  "fmt"
  "sync/atomic"
  "time"
)

var windowSize int       // 窗口大小，单位毫秒
var windowLimit int32    // 一个窗口内，访问次数限制
var startUnixMilli int64 // 窗口的开始时间

var bucketNum int        // 一个窗口被分成几个桶
var arr1bucket []int32   // 所有的桶
var allBucketTotal int32 // 所有的桶的总和

var p1windowTicker *time.Ticker // 定时器，每隔 windowSize/bucketNum 清理一次
var cleanPass int               // 启动的之后需要跳过的清理次数
var bucketIndex int             // 需要清理的桶的下标

func main() {
  // 每25ms发一次a请求
  go func() {
    for {
      try("a")
      time.Sleep(25 * time.Millisecond)
    }
  }()

  // 每50ms，发一次b请求
  go func() {
    for {
      try("b")
      time.Sleep(50 * time.Millisecond)
    }
  }()

  time.Sleep(5 * time.Second)
}

// 初始化
func init() {
  // 1s的窗口限制5次
  windowSize = 1000
  windowLimit = 5
  startUnixMilli = time.Now().UnixMilli()

  bucketNum = 5
  arr1bucket = make([]int32, bucketNum)
  for i := 0; i < bucketNum; i++ {
    arr1bucket[i] = 0
  }
  allBucketTotal = 0

  // 初始化定时器
  p1windowTicker = time.NewTicker(time.Duration(windowSize/bucketNum) * time.Millisecond)
  cleanPass = bucketNum
  bucketIndex = 0

  go clean()
}

// 尝试请求
func try(name string) bool {
  // 判断请求总次数和限制次数
  if allBucketTotal >= windowLimit {
    fmt.Printf("%s, pass:false, arr1bucket:%d, allBucketTotal:%d\r\n", name, arr1bucket, allBucketTotal)
    return false
  }

  // 获取当前时间，计算计数桶的位置
  nowUnixMilli := time.Now().UnixMilli()
  nowBucket := (int(nowUnixMilli-startUnixMilli) / (windowSize / bucketNum)) % bucketNum
  // 原子+1
  atomic.AddInt32(&allBucketTotal, 1)
  atomic.AddInt32(&arr1bucket[nowBucket], 1)

  fmt.Printf("%s, pass:true, arr1bucket[%d]+1:%d, allBucketTotal+1:%d\r\n", name, nowBucket, arr1bucket, allBucketTotal)
  return true
}

// clean 清理方法
func clean() {
  for {
    <-p1windowTicker.C
    // 刚启动窗口还没到最大，不清理
    if cleanPass > 0 {
      cleanPass--
      continue
    }

    // sleep一小会，保证清理的时候，请求全部打下一个桶
    time.Sleep(time.Duration(windowSize/bucketNum/10) * time.Millisecond)
    atomic.AddInt32(&allBucketTotal, -arr1bucket[bucketIndex]) // 总和减去清理桶的计数
    fmt.Printf("clean, free %d, allBucketTotal:%d\r\n", arr1bucket[bucketIndex], allBucketTotal)
    arr1bucket[bucketIndex] = 0                 // 重置清理桶的计数
    bucketIndex = (bucketIndex + 1) % bucketNum // 下标后移

    // 当当前时间和起始时间超过一个窗口大小时，调整起始时间的位置
    nowUnixMilli := time.Now().UnixMilli()
    if int(nowUnixMilli-startUnixMilli) > windowSize {
      startUnixMilli = startUnixMilli + int64(windowSize)
    }
  }
}
