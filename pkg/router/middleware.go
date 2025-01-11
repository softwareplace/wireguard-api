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
		var ctx request.ApiRequestContext

		error_handler.Handler(func() {
			start := time.Now() // Record the start time
			ctx = request.Of(w, r)

			log.Printf("[%s]:: Incoming request: %s %s from %s", ctx.GetSessionId(), r.Method, r.URL.Path, r.RemoteAddr)

			ctx.Next(next)

			duration := time.Since(start)

			var apiKeyId string
			if ctx.AccessContext != nil && ctx.AccessContext.ApiKeyId != "" {
				apiKeyId = ctx.AccessContext.ApiKeyId
			}

			log.Printf("[%s]:: => %s => request processed: %s %s in %v", ctx.GetSessionId(), apiKeyId, r.Method, r.URL.Path, duration)
		}, func(err any) {
			onError(err, w)
		})
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
