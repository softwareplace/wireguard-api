package router

import (
	"github.com/eliasmeireles/wireguard-api/pkg/handlers/request"
	"log"
	"net/http"
	"time"
)

// loggingMiddleware logs each incoming request's method, path, and remote address
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestContext := request.Build(w, r)

		start := time.Now() // Record the start time
		log.Printf("Incoming request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Proceed to the next middleware/handler
		next.ServeHTTP(requestContext.Writer, requestContext.Request)

		// Log the time taken to process the request
		duration := time.Since(start)
		log.Printf("Request processed: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
