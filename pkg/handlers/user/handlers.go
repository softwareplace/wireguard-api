package user

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/security"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
)

type Handler interface {
	UsersRepository() user.UsersRepository
	Login(ctx *api_context.ApiRequestContext[*request.ApiContext])
	CreateUser(ctx *api_context.ApiRequestContext[*request.ApiContext])
	UpdateUser(ctx *api_context.ApiRequestContext[*request.ApiContext])
	JWTService() security.ApiSecurityService[*request.ApiContext]
	Init()
}

type handlerImpl struct {
}

func (h *handlerImpl) UsersRepository() *user.UsersRepository {
	return user.Repository()
}

func Init(api server.ApiRouterHandler[*request.ApiContext]) {
	handler := handlerImpl{}
	api.Post(handler.CreateUser, "user", "POST", "resource:users:create:user")
	api.Put(handler.UpdateUser, "user/:id", "resource:users:update:user")
	api.Put(handler.UpdateUser, "user")
}
