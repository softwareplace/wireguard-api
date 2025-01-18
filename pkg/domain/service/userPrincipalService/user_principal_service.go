package userPrincipalService

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/security/principal"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"sync"
)

type UserPrincipalService struct {
	userRepository user.UsersRepository
}

var (
	once    sync.Once
	service principal.PService[*request.ApiContext]
)

func New() principal.PService[*request.ApiContext] {
	once.Do(func() {
		service = &UserPrincipalService{
			userRepository: *user.Repository(),
		}
	})
	return service
}

func (u *UserPrincipalService) LoadPrincipal(ctx *api_context.ApiRequestContext[*request.ApiContext]) bool {
	userResponse, err := u.userRepository.FindUserBySalt(ctx.Principal.GetSalt())
	if err != nil {
		return false
	}

	ctx.Principal = &request.ApiContext{
		User: userResponse.Parse(),
	}

	return true
}
