package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func main() {
  RunDemo1()
}

func RunDemo1() {
  WriteTxtToFile("./demo/file/demo.txt", "写入文本", false)
  ReadTxtFromFile("./demo/file/demo.txt")
  WriteTxtToFile("./demo/file/demo.txt", "追加文本", false)
  ReadTxtFromFile("./demo/file/demo.txt")
}

//写入文件
//写入时如果本次写入数据长度小于原数据，那么前面的数据会被覆盖，后面的会被保留
func WriteTxtToFile(filePath string, data string, isAppend bool) {
  var pfile *os.File
  var err error
  if isAppend {
    pfile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //只写，不存在就创建
  } else {
    pfile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0666) //追加，不存在就创建
  }
  if err != nil {
    fmt.Printf("打开文件错误：%v\r\n", err)
    return
  }
  defer pfile.Close() //延迟关闭，写在前面防止忘了

  writer := bufio.NewWriter(pfile) //使用带缓存的*Writer
  for index := 0; index < 3; index++ {
    data1 := fmt.Sprintf("%s-%d\r\n", data, index)
    writer.WriteString(data1) //在调用WriterString方法时，内容先写入缓存
  }
  writer.Flush() //调用flush方法，将缓存的数据真正写入到文件中。
  fmt.Printf("写入完成\r\n")
}

//读取文件
func ReadTxtFromFile(filePath string) {
  pfile, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
  if err != nil {
    fmt.Printf("打开文件错误：%v\r\n", err)
    return
  }
  defer pfile.Close()

  reader := bufio.NewReader(pfile) //使用带缓存的*Reader
  for {
    data, err := reader.ReadString('\n') //读取直到第一次遇到\n，返回一个包含已读取的数据和\n的字符串
    if err == io.EOF {
      break
    }
    fmt.Printf("读取一行：%s", data)
  }
  fmt.Printf("读取完成\r\n")
}
