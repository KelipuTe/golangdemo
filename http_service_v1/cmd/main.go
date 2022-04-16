package main

import (
  "demo_golang/http_service_v1"
  "fmt"
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
  fmt.Println("done")
}

func httpApi(p1hsv1 http_service_v1.Service) {
  p1hsv1.RegisteRoute(http.MethodGet, "/user/info", func(p1c *http_service_v1.HTTPContext) { fmt.Println("/user/info1") })
  p1hsv1.RegisteRoute(http.MethodGet, "/user/*", func(p1c *http_service_v1.HTTPContext) { fmt.Println("/user/*") })
  p1hsv1.RegisteRoute(http.MethodGet, "/user/order", func(p1c *http_service_v1.HTTPContext) { fmt.Println("/user/order") })
  p1hsv1.RegisteRoute(http.MethodGet, "/user/info", func(p1c *http_service_v1.HTTPContext) { fmt.Println("/user/info2") })

  p1hsv1.RegisteRoute(http.MethodGet, "/api/text", func(p1c *http_service_v1.HTTPContext) {
    p1c.P1resW.WriteHeader(http.StatusOK)
    _, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /api/text"))
  })

  p1hsv1.RegisteRoute(http.MethodPost, "/api/json", func(p1c *http_service_v1.HTTPContext) {
    reqData := &ApiJson{}
    err := p1c.ReadJson(reqData)
    if nil != err {
      p1c.WriteJson(http.StatusUnprocessableEntity, err.Error())
      return
    }
    reqData.JsonText = "response, http.MethodPost, /api/json"
    p1c.WriteJson(http.StatusOK, reqData)
  })
}
