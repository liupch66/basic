// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"basic-go/webook/internal/repository"
	article2 "basic-go/webook/internal/repository/article"
	"basic-go/webook/internal/repository/cache"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/repository/dao/article"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/service/oauth2/wechat"
	"basic-go/webook/internal/web"
	"basic-go/webook/internal/web/jwt"
	"basic-go/webook/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitUserSvc() service.UserService {
	gormDB := InitTestDB()
	userDAO := dao.NewUserDAO(gormDB)
	cmdable := InitRedis()
	userCache := cache.NewUserCache(cmdable)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	loggerV1 := InitLog()
	userService := service.NewUserService(userRepository, loggerV1)
	return userService
}

func InitCodeSvc() service.CodeService {
	cmdable := InitRedis()
	codeCache := cache.NewCodeCache(cmdable)
	codeRepository := repository.NewCodeRepository(codeCache)
	smsService := ioc.InitSmsService(cmdable)
	codeService := service.NewCodeService(codeRepository, smsService)
	return codeService
}

func InitWechatSvc() wechat.Service {
	loggerV1 := InitLog()
	wechatService := ioc.InitWechatService(loggerV1)
	return wechatService
}

func InitArticleHandler(dao2 article.ArticleDAO) *web.ArticleHandler {
	articleRepository := article2.NewCachedArticleRepository(dao2)
	loggerV1 := InitLog()
	articleService := service.NewArticleService(articleRepository, loggerV1)
	articleHandler := web.NewArticleHandler(articleService, loggerV1)
	return articleHandler
}

func InitWebServer() *gin.Engine {
	loggerV1 := InitLog()
	cmdable := InitRedis()
	handler := jwt.NewRedisJwtHandler(cmdable)
	v := ioc.InitMiddlewares(loggerV1, cmdable, handler)
	userService := InitUserSvc()
	codeService := InitCodeSvc()
	userHandler := web.NewUserHandler(userService, codeService, handler)
	wechatService := InitWechatSvc()
	wechatHandlerConfig := ioc.InitWechatHandlerConfig()
	oAuth2WechatHandler := web.NewOAuth2WechatHandler(wechatService, userService, wechatHandlerConfig, handler)
	gormDB := InitTestDB()
	articleDAO := article.NewGORMArticleDAO(gormDB)
	articleHandler := InitArticleHandler(articleDAO)
	engine := ioc.InitWebServer(v, userHandler, oAuth2WechatHandler, articleHandler)
	return engine
}

// wire.go:

var thirdPS = wire.NewSet(InitTestDB, InitRedis, InitLog)

var userSvcPS = wire.NewSet(dao.NewUserDAO, cache.NewUserCache, repository.NewUserRepository, service.NewUserService)

var codeSvcPS = wire.NewSet(cache.NewCodeCache, repository.NewCodeRepository, ioc.InitSmsService, service.NewCodeService)
