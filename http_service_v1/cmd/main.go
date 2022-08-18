package main

import (
  http_service "demo-golang/http_service_v1"
  "fmt"
  "net/http"
)

type ApiJson struct {
  JsonInt    int    `json:"json_int"`
  JsonString string `json:"json_string"`
  JsonText   string `json:"json_text"`
}

func main() {
  p1hservice := http_service.NewHTTPSrevice(
    "http-service",
    http_service.TestMiddlewareBuilder,
    http_service.TimeCostMiddlewareBuilder,
  )

  httpApi(p1hservice)
  p1hservice.Start("127.0.0.1", "9501")
  fmt.Println("done")
}

// httpApi 注册路由和处理方法
func httpApi(p1hservice http_service.Service) {
  p1hservice.RegisteRoute(http.MethodGet, "/api/test", func(p1c *http_service.HTTPContext) {
    p1c.P1resW.WriteHeader(http.StatusOK)
    _, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /api/test"))
  })

  p1hservice.RegisteRoute(http.MethodPost, "/api/post_json", func(p1c *http_service.HTTPContext) {
    reqData := &ApiJson{}
    err := p1c.ReadJson(reqData)
    if nil != err {
      p1c.WriteJson(http.StatusUnprocessableEntity, err.Error())
      return
    }
    reqData.JsonText = "response, http.MethodPost, /api/json"
    p1c.WriteJson(http.StatusOK, reqData)
  })

  p1hservice.RegisteRoute(http.MethodGet, "/user/info", func(p1c *http_service.HTTPContext) {
    p1c.P1resW.WriteHeader(http.StatusOK)
    _, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/info/1"))
  })

  p1hservice.RegisteRoute(http.MethodGet, "/user/*", func(p1c *http_service.HTTPContext) {
    p1c.P1resW.WriteHeader(http.StatusOK)
    _, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/*"))
  })

  p1hservice.RegisteRoute(http.MethodGet, "/user/order", func(p1c *http_service.HTTPContext) {
    p1c.P1resW.WriteHeader(http.StatusOK)
    _, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/order"))
  })
}
