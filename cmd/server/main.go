package main

import (
	"github.com/softwareplace/goserve/logger"
	"github.com/softwareplace/http-utils/security"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/apiSecretService"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/userPrincipalService"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/user_service"
	"github.com/softwareplace/wireguard-api/pkg/handlers"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
)

var (
	userService          user_service.Service
	securityService      security.ApiSecurityService[*request.ApiContext]
	secreteAccessHandler security.ApiSecretAccessHandler[*request.ApiContext]
	userLoginService     server.LoginService[*request.ApiContext]
)

func factory(appEnv env.ApplicationEnv) {
	userService = user_service.GetService()
	secretKeyProvider := apiSecretService.GetSecretKeyProvider()
	principalService := userPrincipalService.GetUserPrincipalService()
	securityService = security.ApiSecurityServiceBuild(appEnv.ApiSecretAuthorization, principalService)

	secreteAccessHandler = security.ApiSecretAccessHandlerBuild(
		appEnv.ApiSecretKey,
		secretKeyProvider,
		securityService,
	)
	userLoginService = user_service.GetLoginService(securityService)
}

func initializer(apiServer server.ApiRouterHandler[*request.ApiContext]) {
	db.InitMongoDB()
	userService.Init()
	peer.GetService().Load()
	handlers.Init(apiServer)
}

func init() {
	logger.LogSetup()
}

func main() {
	appEnv := env.AppEnv()
	factory(appEnv)

	server.CreateApiRouter[*request.ApiContext]().
		RegisterMiddleware(secreteAccessHandler.HandlerSecretAccess, security.ApiSecretAccessHandlerName).
		RegisterMiddleware(securityService.AuthorizationHandler, security.ApiSecurityHandlerName).
		WithLoginResource(userLoginService).
		EmbeddedServer(initializer).
		StartServer()
}
