package router

import (
	"encoding/json"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/utils/error_handler"
	"log"
	"net/http"
	"time"
)

// rootAppMiddleware logs each incoming request's method, path, and remote address
func rootAppMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Record the start time
		log.Printf("Incoming request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		error_handler.Handler(
			func() {
				requestContext := request.Of(w, r)
				requestContext.Next(next)
			},
			func(err any) {
				onError(err, w)
			})

		duration := time.Since(start)
		log.Printf("request processed: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}

func onError(err any, w http.ResponseWriter) {
	log.Printf("Error processing request: %+v", err)
	w.WriteHeader(http.StatusInternalServerError)
	responseBody := map[string]interface{}{
		"error_handler": "Failed to process request",
		"statusCode":    http.StatusInternalServerError,
		"timestamp":     time.Now().UnixMilli(),
	}

	err = json.NewEncoder(w).Encode(responseBody)

	if err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
