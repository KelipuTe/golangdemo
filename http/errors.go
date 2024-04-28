package http

import "errors"

var (
	ErrParseReqFailed  = errors.New("HTTP 请求报文解析失败")
	ErrParseRespFailed = errors.New("HTTP 响应报文解析失败")
)
