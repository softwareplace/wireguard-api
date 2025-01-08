package user

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/security"
	"github.com/softwareplace/wireguard-api/pkg/router"
	"net/http"
)

type Handler interface {
	UsersRepository() user.UsersRepository
	Login(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	Init()
	JWTService() security.ApiSecurityService
}

type handlerImpl struct{}

func (h *handlerImpl) UsersRepository() user.UsersRepository {
	return user.Repository()
}

func (h *handlerImpl) ApiSecurityService() security.ApiSecurityService {
	return security.GetApiSecurityService()
}

func Init(api router.ApiRouterHandler) {
	handler := handlerImpl{}
	api.PublicRouter(handler.Login, "login", "POST")
	api.Post(handler.CreateUser, "user", "POST", "resource:users:create:user")
	api.Put(handler.UpdateUser, "user/:id", "resource:users:update:user")
	api.Put(handler.UpdateUser, "user")
}
