package api

const (
  TypeRequest  uint8 = iota // 请求
  TypeResponse              // 响应
)

type APIPackage struct {
  // Type 数据包类型，详见 Type 开头的常量
  Type uint8
  // Action 访问的方法
  Action string
  // 数据（经过 json 格式化的结构体）
  Data string
}

type ReqInRegiste struct {
  Name      string   `json:"name"`
  Sli1Route []string `json:"route"`
}
