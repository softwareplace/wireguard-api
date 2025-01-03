package main

import (
	"github.com/eliasmeireles/wireguard-api/pkg/domain/db"
	"github.com/eliasmeireles/wireguard-api/pkg/domain/service/peer"
	"github.com/eliasmeireles/wireguard-api/pkg/handlers"
	"github.com/eliasmeireles/wireguard-api/pkg/router"
)

func main() {
	db.InitMongoDB()
	api := router.GetApiRouterHandler()
	handlers.Init(api)
	peer.GetService().Load()
	api.StartServer()
}
