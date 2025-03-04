package service

import "demo-golang/thirdparty/wire/repo"

type Service struct {
	repo *repo.Repo
}

func NewService(repo *repo.Repo) *Service {
	return &Service{repo: repo}
}
