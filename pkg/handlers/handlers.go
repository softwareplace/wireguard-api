package handlers

import (
	"github.com/eliasmeireles/wireguard-api/pkg/handlers/peer"
	"github.com/eliasmeireles/wireguard-api/pkg/handlers/user"
	"github.com/eliasmeireles/wireguard-api/pkg/router"
)

func Init(api router.ApiRouterHandler) {
	user.Init(api)
	peer.Init(api)
}
