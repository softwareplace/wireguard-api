package peer

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"log"
	"net/http"
)

func (h *handlerImpl) GetAvailablePeer(ctx *request.ApiRequestContext) {
	peer, err, notFound := h.Service().GetAvailablePeer()
	if notFound {
		ctx.Error("No available peer", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("Error getting peer: %v", err)
		ctx.Error("Failed to get peer", http.StatusInternalServerError)
		return
	}

	ctx.Ok(peer)
}
