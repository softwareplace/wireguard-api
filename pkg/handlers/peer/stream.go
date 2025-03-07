package peer

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
)

func (h *handlerImpl) Stream(ctx *api_context.ApiRequestContext[*request.ApiContext]) {
	server.GetRequestBody(ctx, []models.Peer{}, h.save, server.FailedToLoadBody)
}

func (h *handlerImpl) save(ctx *api_context.ApiRequestContext[*request.ApiContext], peers []models.Peer) {
	err := h.Service().Stream(peers)
	if err != nil {
		log.Printf("[%s]:: error saving peers: %v", ctx.GetSessionId(), err)
		ctx.Error("Failed to save peers", http.StatusInternalServerError)
		return
	}
	log.Printf("[%s]:: peers saved successfully", ctx.GetSessionId())
	ctx.Response(map[string]string{"message": "Peers saved successfully"}, http.StatusOK)
}
