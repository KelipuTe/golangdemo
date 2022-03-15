package http_service_v1

import (
  "encoding/json"
  "io"
  "net/http"
)

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

func (c *Context) ReadJson(data interface{}) error {
  reqBody, err := io.ReadAll(c.P1Req.Body)
  if nil != err {
    return err
  }
  return json.Unmarshal(reqBody, data)
}

func (c *Context) WriteJson(status int, data interface{}) error {
  c.ResW.WriteHeader(status)
  resJson, err := json.Marshal(data)
  if nil != err {
    return err
  }
  _, err = c.ResW.Write(resJson)
  if nil != err {
    return err
  }
  return nil
}
