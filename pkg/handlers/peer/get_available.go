package peer

import (
	"encoding/json"
	"net/http"
)

func (h *handlerImpl) GetAvailablePeer(w http.ResponseWriter, r *http.Request) {
	peer, err := h.Service().GetAvailablePeer()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(peer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
