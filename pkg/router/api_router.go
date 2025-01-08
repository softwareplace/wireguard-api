package router

import (
	"github.com/softwareplace/wireguard-api/pkg/auth"
	"net/http"
)

type ApiRouterHandler interface {
	PublicRouter(
		handler func(w http.ResponseWriter, r *http.Request),
		path string,
		method string,
	)
	Add(
		handler func(w http.ResponseWriter, r *http.Request),
		path string,
		method string,
		requiredRoles ...string,
	)
	Get(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
	Post(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
	Put(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
	Delete(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
	Patch(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
	Options(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
	Head(handler func(w http.ResponseWriter,
		r *http.Request),
		path string,
		requiredRoles ...string,
	)
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
	requiredRoles ...string,
) {
	apiRoute.HandleFunc(appEnv.ContextPath+path, handler).Methods(method)
	auth.AddRoles(method+"::"+appEnv.ContextPath+path, requiredRoles...)
}

func (a *apiRouterHandlerImpl) Get(
	handler func(w http.ResponseWriter,
	r *http.Request),
	path string,
	requiredRoles ...string,
) {
	a.Add(handler, path, "GET", requiredRoles...)
}

func (a *apiRouterHandlerImpl) Post(
	handler func(w http.ResponseWriter,
	r *http.Request),
	path string,
	requiredRoles ...string,
) {
	a.Add(handler, path, "POST", requiredRoles...)
}

func (a *apiRouterHandlerImpl) Put(
	handler func(w http.ResponseWriter,
	r *http.Request),
	path string,
	requiredRoles ...string,
) {
	a.Add(handler, path, "PUT", requiredRoles...)
}

func (a *apiRouterHandlerImpl) Delete(
	handler func(w http.ResponseWriter,
	r *http.Request),
	path string,
	requiredRoles ...string,
) {
	a.Add(handler, path, "DELETE", requiredRoles...)
}

func (a *apiRouterHandlerImpl) Patch(
	handler func(w http.ResponseWriter, r *http.Request),
	path string,
	requiredRoles ...string,
) {
	a.Add(handler, path, "PATCH", requiredRoles...)
}

func (a *apiRouterHandlerImpl) Options(
	handler func(w http.ResponseWriter, r *http.Request),
	path string, requiredRoles ...string,
) {
	a.Add(handler, path, "OPTIONS", requiredRoles...)
}

func (a *apiRouterHandlerImpl) Head(
	handler func(w http.ResponseWriter,
	r *http.Request),
	path string,
	requiredRoles ...string,
) {
	a.Add(handler, path, "HEAD", requiredRoles...)
}
