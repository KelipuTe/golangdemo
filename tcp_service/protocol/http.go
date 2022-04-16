package protocol

import (
  "errors"
  "strconv"
  "strings"
)

const (
  HTTP_MSG_MAX_LENGTH = 1048576 // 接收缓冲区最大大小，1024*1024

  HTTP_STATUS_NO_DATA    = 0 // 没有数据
  HTTP_STATUS_NOT_HTTP   = 1 // 没找到\r\n\r\n
  HTTP_STATUS_WRONG_DATA = 2 // 找到\r\n\r\n，但是格式不对
  HTTP_STATUS_TOO_LONG   = 3 // 报文太长
  HTTP_STATUS_ATOI_ERR   = 4 // 请求体长度字符串转长度数字出错
  HTTP_STATUS_NOT_FINISH = 5 // 数据不完整

  STR_X_WWW_FORM_URLENCODED = "application/x-www-form-urlencoded" // Post请求体类型
)

type Http struct {
  ParseStatus int // 解析状态

  HeaderLength int // http报文请求头数据长度
  BodyLength   int // http报文请求体数据长度

  Method  string // http请求方法
  Uri     string // http请求路由
  Version string // http版本

  MapHeader map[string]string // 解析后的请求头
  MapQuery  map[string]string // 解析后的查询参数
  MapBody   map[string]string // 解析后的请求体
}

func (p1this *Http) DataLength(data []byte) (dataLength int, err error) {
  dataLength = 0
  err = nil

  totalLen := len(data)
  if 0 == totalLen {
    p1this.ParseStatus = HTTP_STATUS_NO_DATA
    err = errors.New("HTTP_STATUS_NO_DATA")
    return
  }

  dataStr := string(data)
  // 找到\r\n\r\n的位置，这个位置分隔请求头和请求体
  indexRNRN := strings.Index(dataStr, "\r\n\r\n")
  if -1 == indexRNRN {
    p1this.ParseStatus = HTTP_STATUS_NOT_HTTP
    err = errors.New("HTTP_STATUS_NOT_HTTP")
    return
  }
  if 0 == indexRNRN {
    p1this.ParseStatus = HTTP_STATUS_WRONG_DATA
    err = errors.New("HTTP_STATUS_WRONG_DATA")
    return
  }
  if indexRNRN >= HTTP_MSG_MAX_LENGTH {
    p1this.ParseStatus = HTTP_STATUS_TOO_LONG
    err = errors.New("HTTP_STATUS_TOO_LONG")
    return
  }
  // 请求头长度等于\r\n\r\n的位置下标加上\r\n\r\n的长度
  p1this.HeaderLength = indexRNRN + 4

  // 找到Content-Length的位置
  // 如果有请求体，Content-Length的位置应该在请求头的尾部
  indexCL := strings.Index(dataStr, "Content-Length: ")
  if indexCL > 0 {
    tempStr := dataStr[indexCL+len("Content-Length: "):]
    indexBody := strings.IndexByte(tempStr, '\r')
    // 截取Content-Length的字符串值
    bodyLenStr := tempStr[0:indexBody]
    bodyLenNum, errAtoi := strconv.Atoi(bodyLenStr)
    if nil != errAtoi {
      p1this.ParseStatus = HTTP_STATUS_ATOI_ERR
      err = errors.New("HTTP_STATUS_ATOI_ERR")
      return
    }
    p1this.BodyLength = bodyLenNum
  }

  dataLength = p1this.HeaderLength + p1this.BodyLength
  if totalLen < dataLength {
    // 报文没接收全
    p1this.ParseStatus = HTTP_STATUS_NOT_FINISH
    err = errors.New("HTTP_STATUS_NOT_FINISH")
    return
  }

  return
}

func (p1this *Http) DataDecode(data []byte) (decodeData []byte, err error) {
  dataStr := string(data)
  headerStr := dataStr[0:p1this.HeaderLength]
  bodyStr := dataStr[p1this.HeaderLength:]
  p1this.ParseHeader(headerStr)
  p1this.ParseBody(bodyStr)

  return
}

// 解析请求头
func (p1this *Http) ParseHeader(headerStr string) {
  p1this.MapHeader = make(map[string]string)
  arr1Header := strings.Split(headerStr, "\r\n")
  firstLineStr := arr1Header[0]
  // 第1行
  arr1firstLine := strings.Split(firstLineStr, " ")
  p1this.Method = arr1firstLine[0]
  p1this.Uri = arr1firstLine[1]
  p1this.Version = arr1firstLine[2]
  p1this.ParseQuery(p1this.Uri)
  // 剩下的行
  for _, val := range arr1Header[1:] {
    // 用冒号+空格去切
    arr1kv := strings.Split(val, ": ")
    if 2 == len(arr1kv) {
      p1this.MapHeader[strings.ToLower(arr1kv[0])] = arr1kv[1]
    }
  }
}

// 解析查询参数
func (p1this *Http) ParseQuery(uriStr string) {
  index := strings.Index(uriStr, "?")
  if index > 0 {
    // 有?号
    queryStr := uriStr[index+1:]
    if "" != queryStr {
      // 有查询参数
      p1this.MapQuery = make(map[string]string)
      arr1Query := strings.Split(queryStr, "&")
      for _, val := range arr1Query {
        arr1kv := strings.Split(val, "=")
        if 2 == len(arr1kv) {
          p1this.MapQuery[strings.ToLower(arr1kv[0])] = arr1kv[1]
        }
      }
    }
  }
}

// 解析请求体
func (p1this *Http) ParseBody(bodyStr string) {
  ct, ok := p1this.MapHeader["content-type"]
  if ok {
    switch ct {
    case STR_X_WWW_FORM_URLENCODED:
      p1this.MapBody = make(map[string]string)
      arr1Body := strings.Split(bodyStr, "&")
      for _, val := range arr1Body {
        arr1kv := strings.Split(val, "=")
        if 2 == len(arr1kv) {
          p1this.MapBody[strings.ToLower(arr1kv[0])] = arr1kv[1]
        }
      }
    }
  }
}

func (p1this *Http) DataEncode(data []byte) (encodeData []byte, err error) {
  return
}
