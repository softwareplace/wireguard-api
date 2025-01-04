package router

import (
	"github.com/softwareplace/wireguard-api/pkg/auth"
	"net/http"
)

type ApiRouterHandler interface {
	PublicRouter(handler func(w http.ResponseWriter, r *http.Request), path string, method string)
	Add(handler func(w http.ResponseWriter, r *http.Request), path string, method string)
	Get(handler func(w http.ResponseWriter, r *http.Request), path string)
	Post(handler func(w http.ResponseWriter, r *http.Request), path string)
	Put(handler func(w http.ResponseWriter, r *http.Request), path string)
	Delete(handler func(w http.ResponseWriter, r *http.Request), path string)
	Patch(handler func(w http.ResponseWriter, r *http.Request), path string)
	Options(handler func(w http.ResponseWriter, r *http.Request), path string)
	Head(handler func(w http.ResponseWriter, r *http.Request), path string)
	StartServer()
}

type apiRouterHandlerImpl struct{}

func GetApiRouterHandler() ApiRouterHandler {
	return &apiRouterHandlerImpl{}
}

func (a *apiRouterHandlerImpl) PublicRouter(
	handler func(w http.ResponseWriter, r *http.Request),
	path string,
	method string,
) {
	a.Add(handler, path, method)
	auth.AddOpenPath(method + "::" + appEnv.ContextPath + path)
}

func (a *apiRouterHandlerImpl) Add(
	handler func(w http.ResponseWriter, r *http.Request),
	path string,
	method string,
) {
	apiRoute.HandleFunc(appEnv.ContextPath+path, handler).Methods(method)
}

func (a *apiRouterHandlerImpl) Get(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "GET")
}

func (a *apiRouterHandlerImpl) Post(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "POST")
}

func (a *apiRouterHandlerImpl) Put(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "PUT")
}

func (a *apiRouterHandlerImpl) Delete(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "DELETE")
}

func (a *apiRouterHandlerImpl) Patch(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "PATCH")
}

func (a *apiRouterHandlerImpl) Options(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "OPTIONS")
}

func (a *apiRouterHandlerImpl) Head(handler func(w http.ResponseWriter, r *http.Request), path string) {
	a.Add(handler, path, "HEAD")
}
