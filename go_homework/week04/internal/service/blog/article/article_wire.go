//go:build wireinject
// +build wireinject

package article

import (
  "demo_golang/go_homework/week04/internal/data/model/admin"
  "demo_golang/go_homework/week04/internal/data/model/blog"
  admin_repo "demo_golang/go_homework/week04/internal/data/repository/admin"
  blog_repo "demo_golang/go_homework/week04/internal/data/repository/blog"

  "github.com/google/wire"
)

func NewPublishEvent() *PublishEvent {
  return &PublishEvent{}
}

func NewPublishMission(u *admin.UserModel, a *blog.ArticleModel, e *PublishEvent) *PublishMission {
  return &PublishMission{u, a, e}
}

func InitPublishMission() *PublishMission {
  wire.Build(admin_repo.NewUser, blog_repo.NewArticle, NewPublishEvent, NewPublishMission)
  return &PublishMission{}
}
