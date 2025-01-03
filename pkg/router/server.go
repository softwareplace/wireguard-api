package router

import (
	"github.com/eliasmeireles/wireguard-api/pkg/auth"
	"github.com/eliasmeireles/wireguard-api/pkg/utils/env"
	"github.com/gorilla/mux"
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
