package router

import (
	"github.com/gorilla/mux"
	"github.com/softwareplace/wireguard-api/pkg/auth"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
	"log"
	"net/http"
)

var (
	apiRoute = mux.NewRouter()
	appEnv   = env.AppEnv()
)

func (a *apiRouterHandlerImpl) StartServer() {
	securityHandler := auth.NewApiSecurityHandler()
	securityHandler.InitAPISecretKey()

	apiRoute.Use(loggingMiddleware)
	apiRoute.Use(securityHandler.Middleware)
	apiRoute.Use(auth.AccessValidation)

	log.Printf("Server started at http://localhost:%s%s", appEnv.Port, appEnv.ContextPath)
	log.Fatal(http.ListenAndServe(":"+appEnv.Port, apiRoute))
}
