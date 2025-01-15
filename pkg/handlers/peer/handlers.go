package peer

import (
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
)

type Handler interface {
	GetAvailablePeer(ctx *server.ApiRequestContext)
	Stream(ctx *server.ApiRequestContext)
	Service() peer.Service
}

type handlerImpl struct{}

func GetHandler() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) Service() peer.Service {
	return peer.GetService()
}

func Init(api server.ApiRouterHandler) {
	handler := GetHandler()
	api.Get(handler.GetAvailablePeer, "peers", "resource:peers:get:peer")
	api.Post(handler.Stream, "peers/stream", "resource:peers:stream:peers")
}
