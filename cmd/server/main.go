package main

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/handlers"
	"github.com/softwareplace/wireguard-api/pkg/router"
)

func main() {
	db.InitMongoDB()
	api := router.GetApiRouterHandler()
	handlers.Init(api)
	peer.GetService().Load()
	api.StartServer()
}
