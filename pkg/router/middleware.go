package router

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"log"
	"net/http"
	"time"
)

// loggingMiddleware logs each incoming request's method, path, and remote address
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestContext := request.Of(w, r)

		start := time.Now() // Record the start time
		log.Printf("Incoming request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Proceed to the next middleware/handler
		requestContext.Next(next)

		// Log the time taken to process the request
		duration := time.Since(start)
		log.Printf("request processed: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
