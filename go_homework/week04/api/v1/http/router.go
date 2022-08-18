package http

import (
  admin_mid "demo-golang/go_homework/week04/internal/biz/admin/middleware"
  blog_biz "demo-golang/go_homework/week04/internal/biz/blog/article"

  "github.com/gin-gonic/gin"
)

// 注册 v1 路由
func HttpRegister(p1ge *gin.Engine) {
  p1rg := p1ge.Group("/api/v1/").Use(admin_mid.AuthMiddleWare)
  {
    p1rg.POST("publish_article", blog_biz.PublishArticle)
  }
}
