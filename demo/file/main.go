package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func main() {
  fmt.Println("demo run")
  filePath := "./demo/file/demo.txt"
  writeTxtToFile(filePath, "写入文本", false)
  readTxtFromFile(filePath)
  writeTxtToFile(filePath, "追加文本", false)
  readTxtFromFile(filePath)
  fmt.Println("demo finish")
}

//写入文件
func writeTxtToFile(filePath string, data string, isAppend bool) {
  var p1file *os.File
  var err error

  if isAppend {
    p1file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //只写，不存在就创建
    // 使用只写方式时，如果本次写入数据长度小于原数据，那么原数据前面的会被覆盖，后面的会被保留。
  } else {
    p1file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0666) //追加，不存在就创建
  }
  if err != nil {
    fmt.Printf("打开文件错误：%v\r\n", err)
    return
  }
  defer p1file.Close() //延迟关闭，写在前面防止忘了

  p1writer := bufio.NewWriter(p1file) //使用带缓存的*Writer
  for index := 0; index < 3; index++ {
    data1 := fmt.Sprintf("%s-%d\r\n", data, index)
    p1writer.WriteString(data1) //在调用WriterString方法时，内容先写入缓存
  }
  p1writer.Flush() //调用flush方法，将缓存的数据真正写入到文件中
  fmt.Printf("写入完成\r\n")
}

//读取文件
func readTxtFromFile(filePath string) {
  p1file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
  if err != nil {
    fmt.Printf("打开文件错误：%v\r\n", err)
    return
  }
  defer p1file.Close()

  reader := bufio.NewReader(p1file) //使用带缓存的*Reader
  for {
    data, err := reader.ReadString('\n') //读取直到第一次遇到\n，返回一个包含已读取的数据和\n的字符串
    if err == io.EOF {
      break
    }
    fmt.Printf("读取一行：%s", data)
  }
  fmt.Printf("读取完成\r\n")
}
