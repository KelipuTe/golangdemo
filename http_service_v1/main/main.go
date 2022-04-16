package main

import (
  "demo_golang/http_service_v1"
  "net/http"
)

type ApiJson struct {
  JsonInt    int    `json:"json_int"`
  JsonString string `json:"json_string"`
  JsonText   string `json:"json_text"`
}

func main() {

  p1hsv1 := http_service_v1.NewHTTPSrevice(
    "http-service",
    http_service_v1.Test1MiddlewareBuilder,
    http_service_v1.Test2MiddlewareBuilder,
    http_service_v1.TimeCostMiddlewareBuilder,
  )

  httpApi(p1hsv1)
  p1hsv1.Start("127.0.0.1", "9501")
}

// 注册 http 路由和对应的处理方法
func httpApi(p1hsv1 http_service_v1.Service) {
  p1hsv1.RegisteRoute("GET", "/api/text", func(p1c *http_service_v1.HTTPContext) {
    p1c.P1resW.WriteHeader(http.StatusOK)
    _, _ = p1c.P1resW.Write([]byte("GET /api/text response"))
  })

  p1hsv1.RegisteRoute("POST", "/api/json", func(p1c *http_service_v1.HTTPContext) {
    reqData := &ApiJson{}
    err := p1c.ReadJson(reqData)
    if nil != err {
      p1c.WriteJson(http.StatusUnprocessableEntity, err.Error())
      return
    }
    reqData.JsonText = "POST /api/json response"
    p1c.WriteJson(http.StatusOK, reqData)
  })
}
