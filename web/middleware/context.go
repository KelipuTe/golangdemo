package middleware

import "net/http"

type HTTPContext struct {
	I9writer  http.ResponseWriter
	P7request *http.Request

	// RespData 暂存请求数据
	// 因为 http.Request.Body 是流，只能读一次（和 linux c 的 recvFrom() 类似）
	// 如果等到应用层再调用，那么在中间件里面记录请求日志或者进行预处理就无法实现
	// 这里的方案是，在所有的处理流程开始前就读取然后存下来，如果有需要可以再造一个流放回去
	ReqBody []byte
	// RespData 暂存响应 http status code
	RespStatusCode int
	// RespData 暂存响应数据
	// 因为 http.ResponseWriter.Write 是流，只能写一次（和 linux c 的 write() 类似）
	// 如果在应用层调用了，那么在中间件里面记录响应日志或者追加数据就无法实现
	// 这里的方案是，等到所有的处理流程都结束了，再调用 http.ResponseWriter.Write
	RespData []byte
}
