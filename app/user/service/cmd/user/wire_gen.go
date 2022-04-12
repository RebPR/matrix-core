// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Cube-v2/cube-core/app/user/service/internal/biz"
	"github.com/Cube-v2/cube-core/app/user/service/internal/conf"
	"github.com/Cube-v2/cube-core/app/user/service/internal/data"
	"github.com/Cube-v2/cube-core/app/user/service/internal/server"
	"github.com/Cube-v2/cube-core/app/user/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	cmdable := data.NewRedis(confData, logger)
	txCode := data.NewPhoneCode(confData)
	goMail := data.NewGoMail(confData)
	dataData, cleanup, err := data.NewData(db, cmdable, txCode, goMail, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	authRepo := data.NewAuthRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(auth, authRepo, userRepo, logger)
	userService := service.NewUserService(userUseCase, authUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
