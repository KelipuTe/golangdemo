package middleware

import (
  "log"

  "github.com/gin-gonic/gin"
)

// 鉴权中间件
func AuthMiddleWare(p1gc *gin.Context) {
  log.Println("admin AuthMiddleWare")
  // 一通操作拿到请求的用户id
  p1gc.Set("user_id", 1)
  p1gc.Next()
}
