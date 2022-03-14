package http_service_v1

import "net/http"

type Context struct {
  ResW  http.ResponseWriter
  P1Req *http.Request
}

func NewContext(resW http.ResponseWriter, p1Req *http.Request) *Context {
  return &Context{
    ResW:  resW,
    P1Req: p1Req,
  }
}
