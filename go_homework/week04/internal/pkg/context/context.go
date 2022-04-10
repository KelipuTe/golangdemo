package context

import (
  "context"

  "github.com/gin-gonic/gin"

  pkg_log "demo_golang/go_homework/week04/internal/pkg/log"
)

type BlogContextKey string

type BlogContextContext struct {
  FlowId string // 流程 id
}

const BLOG_CONTEXT_KEY BlogContextKey = "blog_context"

func NewBlogContext() *BlogContextContext {
  return &BlogContextContext{
    FlowId: "",
  }
}

func MakeBlogContext(p1gc *gin.Context) context.Context {
  p1bc := NewBlogContext()
  p1bc.FlowId = pkg_log.MakeFlowId()
  return context.WithValue(p1gc, BLOG_CONTEXT_KEY, p1bc)
}

func GetBlogContext(c context.Context) *BlogContextContext {
  value := c.Value(BLOG_CONTEXT_KEY)
  if nil == value {
    return NewBlogContext()
  }
  p1bc, ok := value.(*BlogContextContext)
  if ok {
    return p1bc
  } else {
    return NewBlogContext()
  }
}
