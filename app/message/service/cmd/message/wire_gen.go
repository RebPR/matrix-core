// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/message/service/internal/biz"
	"github.com/the-zion/matrix-core/app/message/service/internal/conf"
	"github.com/the-zion/matrix-core/app/message/service/internal/data"
	"github.com/the-zion/matrix-core/app/message/service/internal/server"
	"github.com/the-zion/matrix-core/app/message/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logLogger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	cmdable := data.NewRedis(confData)
	userClient := data.NewUserServiceClient(registry)
	creationClient := data.NewCreationServiceClient(registry)
	commentClient := data.NewCommentServiceClient(registry)
	achievementClient := data.NewAchievementServiceClient(registry)
	jwt := data.NewJwtClient(confData)
	cosUser := data.NewCosUserClient(confData)
	cosCreation := data.NewCosCreationClient(confData)
	cosComment := data.NewCosCommentClient(confData)
	txCode := data.NewPhoneCode(confData)
	goMail := data.NewGoMail(confData)
	dataData, cleanup2, err := data.NewData(db, cmdable, userClient, creationClient, commentClient, achievementClient, jwt, cosUser, cosCreation, cosComment, txCode, goMail, logLogger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logLogger)
	messageRepo := data.NewMessageRepo(dataData, logLogger)
	transaction := data.NewTransaction(dataData)
	bizJwt := data.NewJwt(dataData)
	userUseCase := biz.NewUserUseCase(userRepo, messageRepo, transaction, bizJwt, logLogger)
	creationRepo := data.NewCreationRepo(dataData, logLogger)
	creationUseCase := biz.NewCreationUseCase(creationRepo, messageRepo, transaction, bizJwt, logLogger)
	achievementRepo := data.NewAchievementRepo(dataData, logLogger)
	commentRepo := data.NewCommentRepo(dataData, logLogger)
	recovery := data.NewRecovery(dataData)
	achievementUseCase := biz.NewAchievementUseCase(achievementRepo, creationRepo, commentRepo, recovery, logLogger)
	commentUseCase := biz.NewCommentUseCase(commentRepo, messageRepo, transaction, bizJwt, logLogger)
	messageUseCase := biz.NewMessageUseCase(messageRepo, recovery, logLogger)
	messageService := service.NewMessageService(userUseCase, creationUseCase, achievementUseCase, commentUseCase, messageUseCase, logLogger)
	httpServer := server.NewHTTPServer(confServer, messageService, logLogger)
	grpcServer := server.NewGRPCServer(confServer, messageService, logLogger)
	codeMqConsumerServer := server.NewCodeMqConsumerServer(confServer, messageService, logLogger)
	profileMqConsumerServer := server.NewProfileMqConsumerServer(confServer, messageService, logLogger)
	pictureMqConsumerServer := server.NewPictureMqConsumerServer(confServer, messageService, logLogger)
	followMqConsumerServer := server.NewFollowMqConsumerServer(confServer, messageService, logLogger)
	articleReviewMqConsumerServer := server.NewArticleReviewMqConsumerServer(confServer, messageService, logLogger)
	articleMqConsumerServer := server.NewArticleMqConsumerServer(confServer, messageService, logLogger)
	talkReviewMqConsumerServer := server.NewTalkReviewMqConsumerServer(confServer, messageService, logLogger)
	talkMqConsumerServer := server.NewTalkMqConsumerServer(confServer, messageService, logLogger)
	columnReviewMqConsumerServer := server.NewColumnReviewMqConsumerServer(confServer, messageService, logLogger)
	columnMqConsumerServer := server.NewColumnMqConsumerServer(confServer, messageService, logLogger)
	achievementMqConsumerServer := server.NewAchievementMqConsumerServer(confServer, messageService, logLogger)
	commentReviewMqConsumerServer := server.NewCommentReviewMqConsumerServer(confServer, messageService, logLogger)
	commentMqConsumerServer := server.NewCommentMqConsumerServer(confServer, messageService, logLogger)
	collectionsReviewMqConsumerServer := server.NewCollectionsReviewMqConsumerServer(confServer, messageService, logLogger)
	collectionsMqConsumerServer := server.NewCollectionsMqConsumerServer(confServer, messageService, logLogger)
	kratosApp := newApp(registry, httpServer, grpcServer, codeMqConsumerServer, profileMqConsumerServer, pictureMqConsumerServer, followMqConsumerServer, articleReviewMqConsumerServer, articleMqConsumerServer, talkReviewMqConsumerServer, talkMqConsumerServer, columnReviewMqConsumerServer, columnMqConsumerServer, achievementMqConsumerServer, commentReviewMqConsumerServer, commentMqConsumerServer, collectionsReviewMqConsumerServer, collectionsMqConsumerServer)
	return kratosApp, func() {
		cleanup2()
	}, nil
}
