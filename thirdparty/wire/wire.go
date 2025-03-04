//go:build wireinject

package wire

import (
	"demo-golang/thirdparty/wire/repo"
	"demo-golang/thirdparty/wire/repo/dao"
	"demo-golang/thirdparty/wire/service"
	"github.com/google/wire"
)

func InitService() *service.Service {
	wire.Build(dao.NewDao, repo.NewRepo, service.NewService)
	return &service.Service{}
}
