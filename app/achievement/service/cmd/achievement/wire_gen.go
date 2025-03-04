// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/achievement/service/internal/biz"
	"github.com/the-zion/matrix-core/app/achievement/service/internal/conf"
	"github.com/the-zion/matrix-core/app/achievement/service/internal/data"
	"github.com/the-zion/matrix-core/app/achievement/service/internal/server"
	"github.com/the-zion/matrix-core/app/achievement/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logLogger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	cmdable := data.NewRedis(confData)
	achievementMqPro := data.NewRocketmqAchievementProducer(confData)
	dataData, cleanup2, err := data.NewData(db, cmdable, achievementMqPro, logLogger)
	if err != nil {
		return nil, nil, err
	}
	achievementRepo := data.NewAchievementRepo(dataData, logLogger)
	recovery := data.NewRecovery(dataData)
	transaction := data.NewTransaction(dataData)
	achievementUseCase := biz.NewAchievementUseCase(achievementRepo, recovery, transaction, logLogger)
	achievementService := service.NewAchievementService(achievementUseCase, logLogger)
	httpServer := server.NewHTTPServer(confServer, achievementService, logLogger)
	grpcServer := server.NewGRPCServer(confServer, achievementService, logLogger)
	kratosApp := newApp(registry, httpServer, grpcServer)
	return kratosApp, func() {
		cleanup2()
	}, nil
}
