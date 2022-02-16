package http

import "fmt"

// 响应结构体
type Response struct {
  MapHeader map[string]string // 响应头
}

var (
  MapRespCodeAndStr map[int]string // http状态码和文案
)

func init() {
  MapRespCodeAndStr = make(map[int]string)
  MapRespCodeAndStr[200] = "OK"
  MapRespCodeAndStr[400] = "Bad Request"
  MapRespCodeAndStr[401] = "Unauthorized"
  MapRespCodeAndStr[403] = "Forbidden"
  MapRespCodeAndStr[404] = "Not Found"
  MapRespCodeAndStr[405] = "Method Not Allowed"
  MapRespCodeAndStr[406] = "Not Acceptable"
  MapRespCodeAndStr[500] = "Internal Server Error"
  MapRespCodeAndStr[501] = "Not Implemented"
  MapRespCodeAndStr[502] = "Bad Gateway"
}

// 手动初始化
func (p1this *Response) HandInit() {
  p1this.MapHeader = make(map[string]string)
}

// 设置响应头
func (p1this *Response) SetHeader(key string, val string) {
  p1this.MapHeader[key] = val
}

// 构造响应数据
func (p1this *Response) MakeData(httpCode int, data string) (respStr string) {
  respStr = fmt.Sprintf("HTTP/1.1 %d %v\r\n", httpCode, MapRespCodeAndStr[httpCode])

  _, ok := p1this.MapHeader["Content-Type"]
  if !ok {
    p1this.MapHeader["Content-Type"] = "text/html;charset=utf8"
  }
  for key, val := range p1this.MapHeader {
    respStr += fmt.Sprintf("%v: %v\r\n", key, val)
  }

  dataLen := len(data)
  respStr += fmt.Sprintf("Content-Length: %v\r\n", dataLen)
  respStr += "\r\n"
  respStr += data

  return
}
