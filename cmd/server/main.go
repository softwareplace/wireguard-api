package main

import (
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/auth"
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/user"
	"github.com/softwareplace/wireguard-api/pkg/handlers"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
)

func main() {
	db.InitMongoDB()
	api := server.New()
	api.Router().Use(request.ContextBuilder)
	handler := auth.NewApiSecurityHandler()
	api.Router().Use(handler.Middleware)
	api.Router().Use(auth.AccessValidation)
	handlers.Init(api)
	peer.GetService().Load()
	user.GetService().Init()
	api.StartServer()
}
