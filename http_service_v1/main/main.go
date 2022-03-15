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
  // 创建 HTTP 服务，并指定中间件
  httpService := http_service_v1.NewHTTPSrevice(
    "http-service",
    http_service_v1.TestMiddlewareBuilder,
    http_service_v1.TimeCostMiddlewareBuilder,
  )

  // 注册路由和对应的处理方法
  httpService.RegisteRoute("GET", "/api/text", func(c *http_service_v1.Context) {
    c.ResW.WriteHeader(http.StatusOK)
    _, _ = c.ResW.Write([]byte("GET /api/text response"))
  })
  httpService.RegisteRoute("POST", "/api/json", func(c *http_service_v1.Context) {
    reqData := &ApiJson{}
    err := c.ReadJson(reqData)
    if nil != err {
      c.WriteJson(http.StatusUnprocessableEntity, err.Error())
      return
    }
    reqData.JsonText = "POST /api/json response"
    c.WriteJson(http.StatusOK, reqData)
  })

  // 启动服务
  httpService.Start("127.0.0.1", "9501")
}
