package http

const (
	connPoolNumMax = 1024 //最大tcp连接数
)

const (
	readBufferMaxLen = 1048576 // 1048576 == 2^20 == 1MB。
)

const (
	VersionX1Y1 = "HTTP/1.1"

	MethodGet  = "GET"
	MethodPost = "POST"
)

const (
	StatusSwitchingProtocols = 101

	StatusOK = 200
)

var (
	statusText = map[int]string{
		StatusSwitchingProtocols: "Switching Protocols",

		StatusOK: "OK",
	}
)
