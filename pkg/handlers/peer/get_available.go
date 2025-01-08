package peer

import (
	"encoding/json"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"log"
	"net/http"
)

func (h *handlerImpl) GetAvailablePeer(w http.ResponseWriter, r *http.Request) {
	peer, err, notFound := h.Service().GetAvailablePeer()
	if notFound {
		shared.MakeErrorResponse(w, "No available peer", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("Error getting peer: %v", err)
		shared.MakeErrorResponse(w, "Failed to get peer", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(peer); err != nil {
		log.Printf("Error getting peer: %v", err)
		shared.MakeErrorResponse(w, "Failed to get peer", http.StatusInternalServerError)
	}
}
