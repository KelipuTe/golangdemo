package protocol

// Request 请求数据结构
type Request struct {
	ServiceName   string
	FuncName      string
	MetaData      map[string]string // 元数据
	SerializeCode uint8             // 方法入参的序列化方式
	FuncInput     []byte            // 方法入参
}

// Response 响应数据结构
type Response struct {
	Error         error  // 异常
	SerializeCode uint8  // 方法出参的序列化方式
	FuncOutput    []byte // 方法出参
}
