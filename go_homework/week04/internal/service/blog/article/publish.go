package article

import (
  "context"
  "demo_golang/go_homework/week04/internal/data/model/admin"
  "demo_golang/go_homework/week04/internal/data/model/blog"
  blog_repo "demo-golang/go_homework/week04/internal/data/repository/blog"
  "log"

  "github.com/pkg/errors"
)

type ArticleBizModel struct {
  Content string
}

type PublishEvent struct {
  UserId    int64
  ArticleId int64
}

type PublishMission struct {
  P1Writer  *admin.UserModel
  P1Article *blog.ArticleModel
  P1Event   *PublishEvent
}

// 发布
func PublishArticle(c context.Context, userId int64, p1article *ArticleBizModel) (int64, error) {
  log.Printf("PublishArticle, userId=%d, Content=%s", userId, p1article.Content)
  p1mission := InitPublishMission()
  articleId, err := p1mission.SaveArticle(c, userId, p1article)
  if nil != err {
    return 0, err
  }
  p1mission.PushEvent(c)
  return articleId, nil
}

// 存储
func (p1 *PublishMission) SaveArticle(c context.Context, userId int64, p1article *ArticleBizModel) (int64, error) {
  log.Printf("SaveArticle, userId=%d, Content=%s", userId, p1article.Content)

  p1.P1Writer.Id = userId
  p1.P1Article.UserId = userId
  p1.P1Event.UserId = userId

  p1.P1Article.Content = p1article.Content
  articleId, err := blog_repo.InsertArticle(c, p1.P1Article)
  if nil != err {
    return 0, errors.WithStack(err)
  }
  p1.P1Article.Id = articleId
  p1.P1Event.ArticleId = articleId

  go p1.PushEvent(c)

  return articleId, nil
}

// 事件
func (p1 *PublishMission) PushEvent(c context.Context) {
  log.Printf("PushEvent, userId=%d, articleId=%d", p1.P1Event.UserId, p1.P1Event.ArticleId)
}
