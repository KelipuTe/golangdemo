package api

const (
  TypeRequest  uint8 = iota // 请求
  TypeResponse              // 响应
)

// 自定义的交互数据结构
type APIPackage struct {
  // 数据结构的 ID
  Id string
  // Type 数据包类型，详见 Type 开头的常量
  Type uint8
  // Action 访问的方法
  Action string
  // 数据（经过 json 格式化的结构体）
  Data string
}

//
type ReqInRegisteServiceProvider struct {
  Name      string   `json:"name"`
  Sli1Route []string `json:"route"`
}
