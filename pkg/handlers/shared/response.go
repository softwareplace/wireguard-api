package shared

import (
	"encoding/json"
	"net/http"
	"time"
)

func MakeErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       message,
		"status_code": status,
		"timestamp":   time.Now().UnixMilli(),
	})
}

func MakeResponse(w http.ResponseWriter, body map[string]interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
