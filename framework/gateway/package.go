package gateway

const (
	PackageTypeReq  = 1
	PackageTypeResp = 2
)

// Package 网关和内部服务之间通信的数据结构
type Package struct {
	From    string `json:"from"`    //数据包是哪个连接发来的
	To      string `json:"to"`      //数据包是发给哪个连接的
	Type    int    `json:"type"`    //请求还是响应
	Service string `json:"service"` //外部请求想请求哪个服务
	Uri     string `json:"uri"`     //外部请求想请求哪个接口
	Data    string `json:"data"`    //外部请求的数据
}
