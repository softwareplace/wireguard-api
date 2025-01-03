package user

import (
	"github.com/eliasmeireles/wireguard-api/pkg/domain/repository/user"
	"github.com/eliasmeireles/wireguard-api/pkg/domain/service/security"
	"github.com/eliasmeireles/wireguard-api/pkg/router"
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
	api.PublicRouter(handler.CreateUser, "user", "POST")
	api.Put(handler.UpdateUser, "user")
}
