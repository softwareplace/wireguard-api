package peer

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/service/peer"
	"github.com/softwareplace/wireguard-api/pkg/router"
	"net/http"
)

type Handler interface {
	GetAvailablePeer(w http.ResponseWriter, r *http.Request)
	Service() peer.Service
	Stream(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct{}

func GetHandler() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) Service() peer.Service {
	return peer.GetService()
}

func Init(api router.ApiRouterHandler) {
	handler := GetHandler()
	api.Get(handler.GetAvailablePeer, "peers", "resource:peers:get:peer")
	api.Post(handler.Stream, "peers/stream", "resource:peers:stream:peers")
}
