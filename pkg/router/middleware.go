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
		error_handler.Handler(func() {
			start := time.Now() // Record the start time
			ctx := request.Of(w, r, "MIDDLEWARE/ROOT_APP")

			log.Printf("[%s]:: Incoming request: %s %s from %s", ctx.GetSessionId(), r.Method, r.URL.Path, r.RemoteAddr)

			ctx.Next(next)

			duration := time.Since(start)
			log.Printf("[%s]:: => %s/%s => request processed: %s %s in %v",
				ctx.GetSessionId(),
				ctx.GetAccessApiKeyId(),
				ctx.GetAccessId(),
				r.Method,
				r.URL.Path,
				duration,
			)

			error_handler.Handler(ctx.Flush, func(err any) {
				log.Printf("[%s]:: Error flushing response: %v", ctx.GetSessionId(), err)
			})
		}, func(err any) {
			onError(err, w)
		})
	})
}

func onError(err any, w http.ResponseWriter) {
	log.Printf("Error processing request: %+v", err)

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	responseBody := map[string]interface{}{
		"message":    "Failed to process request",
		"statusCode": http.StatusInternalServerError,
		"timestamp":  time.Now().UnixMilli(),
	}

	err = json.NewEncoder(w).Encode(responseBody)

	if err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
