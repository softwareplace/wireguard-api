package main

import (
	auth "github.com/softwareplace/http-utils/oauth"
	"github.com/softwareplace/http-utils/security"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/api_secret_service"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/user_service"
	"github.com/softwareplace/wireguard-api/pkg/handlers"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/handlers/user/user_handler"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
)

func main() {
	appEnv := env.AppEnv()
	db.InitMongoDB()
	service := user_service.GetService()
	userAuthenticationUserHandler := user_handler.GetAuthenticationUserHandler(&service)
	api := server.New[*request.ApiContext](
		request.ContextBuilder,
		userAuthenticationUserHandler.Handler,
	)

	secretService := api_secret_service.GetService()
	securityService := security.GetApiSecurityService[*request.ApiContext](appEnv.ApiSecretAuthorization)
	auth.Handler(appEnv.ApiSecretKey, secretService.GetKey, &securityService, api)
	handlers.Init(api)
	peer.GetService().Load()
	service.Init()
	api.StartServer()
}
