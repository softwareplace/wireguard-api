package peer

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
)

type Handler interface {
	GetAvailablePeer(ctx *api_context.ApiRequestContext[*request.ApiContext])
	Stream(ctx *api_context.ApiRequestContext[*request.ApiContext])
	Service() peer.Service
}

type handlerImpl struct{}

func GetHandler() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) Service() peer.Service {
	return peer.GetService()
}

func Init(api server.ApiRouterHandler[*request.ApiContext]) {
	handler := GetHandler()
	api.Get(handler.GetAvailablePeer, "peers", "resource:peers:get:peer")
	api.Post(handler.Stream, "peers/stream", "resource:peers:stream:peers")
}
