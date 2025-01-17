package user_handler

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/user_service"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
)

type AuthenticationUserHandler interface {
	Handler(ctx *api_context.ApiRequestContext[*request.ApiContext]) bool
}

type _AuthenticationUserHandlerImpl struct {
	service *user_service.Service
}

func GetAuthenticationUserHandler(service *user_service.Service) AuthenticationUserHandler {
	return &_AuthenticationUserHandlerImpl{
		service: service,
	}
}
func (h *_AuthenticationUserHandlerImpl) Handler(ctx *api_context.ApiRequestContext[*request.ApiContext]) bool {
	rolesLoader := (*h.service).LoadUserRoles
	ctx.AccessRolesLoader = &rolesLoader
	return true
}
