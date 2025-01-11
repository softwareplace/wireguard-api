package peer

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"log"
	"net/http"
)

func (h *handlerImpl) GetAvailablePeer(ctx *request.ApiRequestContext) {
	peer, err, notFound := h.Service().GetAvailablePeer()
	if notFound {
		log.Printf("[%s]:: no peer available: %v", ctx.GetSessionId(), err)
		ctx.Error("No available peer", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("[%s]:: Error getting peer: %v", ctx.GetSessionId(), err)
		ctx.Error("Failed to get peer", http.StatusInternalServerError)
		return
	}

	ctx.Ok(peer)
}
