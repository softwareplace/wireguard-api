package router

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
)

type ApiContextHandler func(ctx *request.ApiRequestContext)

type ApiRouterHandler interface {
	PublicRouter(handler ApiContextHandler, path string, method string)
	Add(handler ApiContextHandler, path string, method string, requiredRoles ...string)
	Get(handler ApiContextHandler, path string, requiredRoles ...string)
	Post(handler ApiContextHandler, path string, requiredRoles ...string)
	Put(handler ApiContextHandler, path string, requiredRoles ...string)
	Delete(handler ApiContextHandler, path string, requiredRoles ...string)
	Patch(handler ApiContextHandler, path string, requiredRoles ...string)
	Options(handler ApiContextHandler, path string, requiredRoles ...string)
	Head(handler ApiContextHandler, path string, requiredRoles ...string)
	StartServer()
}

type apiRouterHandlerImpl struct{}

func GetApiRouterHandler() ApiRouterHandler {
	return &apiRouterHandlerImpl{}
}
