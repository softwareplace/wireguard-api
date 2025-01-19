package user_service

import (
	"github.com/softwareplace/http-utils/security"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"sync"
	"time"
)

type loginServiceImpl struct {
	securityService security.ApiSecurityService[*request.ApiContext]
	repository      user.UsersRepository
}

func (l *loginServiceImpl) SecurityService() security.ApiSecurityService[*request.ApiContext] {
	return l.securityService
}

var (
	loginServiceOnce     sync.Once
	loginServiceInstance server.LoginService[*request.ApiContext]
)

func GetLoginService(securityService security.ApiSecurityService[*request.ApiContext]) server.LoginService[*request.ApiContext] {
	loginServiceOnce.Do(func() {
		loginServiceInstance = &loginServiceImpl{
			securityService: securityService,
			repository:      user.Repository(),
		}
	})

	return loginServiceInstance
}

func (l *loginServiceImpl) Login(user server.LoginEntryData) (*request.ApiContext, error) {
	response, err := l.repository.FindUserByUsernameOrEmail(user.Username, user.Email)
	if err != nil {
		return nil, err
	}
	return &request.ApiContext{
		User: response.Parse(),
	}, nil
}

func (l *loginServiceImpl) TokenDuration() time.Duration {
	return time.Minute * 15
}
