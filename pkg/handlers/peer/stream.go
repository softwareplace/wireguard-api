package peer

import (
	"encoding/json"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
)

func (h *handlerImpl) Stream(w http.ResponseWriter, r *http.Request) {
	var peers []models.Peer
	err := json.NewDecoder(r.Body).Decode(&peers)

	if err != nil {
		shared.MakeErrorResponse(w, "Failed to decode json body", http.StatusInternalServerError)
		return

	}
	err = h.Service().Stream(peers)
	if err != nil {
		log.Printf("Error saving peers: %v", err)
		shared.MakeErrorResponse(w, "Failed to save peers", http.StatusInternalServerError)
		return
	}
	log.Printf("Peers saved successfully")
	shared.MakeErrorResponse(w, "Peers saved successfully", http.StatusCreated)
}
