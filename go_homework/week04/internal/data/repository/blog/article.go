package blog

import (
	"context"
	"demo_golang/go_homework/week04/internal/data/model/blog"
	"log"
	"time"
)

func NewArticle() *blog.ArticleModel {
  return &blog.ArticleModel{}
}

func InsertArticle(c context.Context, p1orm *blog.ArticleModel) (int64, error) {
  // 这里一般需要从上下文获取源数据决定存到哪里去
  id := time.Now().Unix()
  log.Printf("insert article into mysql, id=%d, content=%s", id, p1orm.Content)
  return id, nil
}
