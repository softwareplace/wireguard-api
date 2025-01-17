package main

import (
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/user_service"
	"github.com/softwareplace/wireguard-api/pkg/handlers"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/handlers/user/user_handler"
)

func main() {
	db.InitMongoDB()
	api := server.New[*request.ApiContext]()
	api.Use(request.ContextBuilder, "API/CONTEXT/INITIALIZER")

	service := user_service.GetService()
	userAuthenticationUserHandler := user_handler.GetAuthenticationUserHandler(&service)
	api.Use(userAuthenticationUserHandler.Handler, "MIDDLEWARE/AUTHENTICATION_USER")

	handlers.Init(&api)
	peer.GetService().Load()
	service.Init()
	api.StartServer()
}
