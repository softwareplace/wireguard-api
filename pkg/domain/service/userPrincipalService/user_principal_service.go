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
	userPrincipalServiceOnce     sync.Once
	userPrincipalServiceInstance principal.PService[*request.ApiContext]
)

func GetUserPrincipalService() principal.PService[*request.ApiContext] {
	userPrincipalServiceOnce.Do(func() {
		userPrincipalServiceInstance = &UserPrincipalService{
			userRepository: user.Repository(),
		}
	})
	return userPrincipalServiceInstance
}

func (u *UserPrincipalService) LoadPrincipal(ctx *api_context.ApiRequestContext[*request.ApiContext]) bool {
	userResponse, err := u.userRepository.FindUserBySalt(ctx.AccessId)
	if err != nil {
		return false
	}

	context := request.NewApiContext(userResponse.Parse())
	ctx.Principal = &context
	return true
}
