package router

import (
	"github.com/eliasmeireles/wireguard-api/pkg/auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	basePath = "/api/private-network/v1/"
	apiRoute = mux.NewRouter()
)

func (a *apiRouterHandlerImpl) StartServer() {
	securityHandler := auth.NewApiSecurityHandler()
	securityHandler.InitAPISecretKey()

	apiRoute.Use(loggingMiddleware)
	apiRoute.Use(securityHandler.Middleware)
	apiRoute.Use(auth.AccessValidation)

	port := "8080" // Default port
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if contextPath := os.Getenv("CONTEXT_PATH"); contextPath != "" {
		basePath = contextPath
	}

	log.Printf("Server started at http://localhost:%s%s", port, basePath)
	log.Fatal(http.ListenAndServe(":"+port, apiRoute))
}
