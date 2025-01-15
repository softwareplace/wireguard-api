package handlers

import (
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/handlers/peer"
	"github.com/softwareplace/wireguard-api/pkg/handlers/user"
)

func Init(api server.ApiRouterHandler) {
	user.Init(api)
	peer.Init(api)
}
