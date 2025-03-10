// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"demo-golang/thirdparty/wire/repo"
	"demo-golang/thirdparty/wire/repo/dao"
	"demo-golang/thirdparty/wire/service"
)

// Injectors from wire.go:

func InitService() *service.Service {
	daoDao := dao.NewDao()
	repoRepo := repo.NewRepo(daoDao)
	serviceService := service.NewService(repoRepo)
	return serviceService
}
