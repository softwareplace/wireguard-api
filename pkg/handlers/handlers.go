package handlers

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/peer"
	"github.com/softwareplace/wireguard-api/pkg/handlers/user"
	"github.com/softwareplace/wireguard-api/pkg/router"
)

func Init(api router.ApiRouterHandler) {
	user.Init(api)
	peer.Init(api)
}
