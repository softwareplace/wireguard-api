package main

import (
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

func main() {
	appEnv := env.AppEnv()
	db.InitMongoDB()

	userService := user_service.GetService()
	principalService := userPrincipalService.New()
	secretService := apiSecretService.GetService()

	securityService := security.ApiSecurityServiceBuild(appEnv.ApiSecretAuthorization, &principalService)
	secreteAccessHandler := security.ApiSecretAccessHandlerBuild(
		appEnv.ApiSecretKey,
		secretService.GetKey,
		securityService,
	)

	userLoginService := user_service.New(securityService)

	api := server.CreateApiRouter[*request.ApiContext]().
		RegisterMiddleware(secreteAccessHandler.HandlerSecretAccess, security.ApiSecretAccessHandlerName).
		RegisterMiddleware(securityService.AuthorizationHandler, security.ApiSecurityHandlerName).
		WithLoginResource(&userLoginService)

	handlers.Init(api)
	userService.Init()
	peer.GetService().Load()
	api.StartServer()
}
