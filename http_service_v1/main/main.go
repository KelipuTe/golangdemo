package main

import (
  "demo_golang/http_service_v1"
  "fmt"
)

func main() {
  httpService := http_service_v1.NewHTTPSrevice("http-service")
  httpService.Route("GET", "/api/test", func(c *http_service_v1.Context) {
    fmt.Println("/api/test")
  })
  httpService.Route("POST", "/api/test2", func(c *http_service_v1.Context) {
    fmt.Println("/api/test2")
  })
  httpService.Start("127.0.0.1", "9501")
}
