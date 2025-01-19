package apiSecretService

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/security"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/api_secret"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"sync"
)

type apiSecretKeyProviderImpl struct {
	repository api_secret.ApiSecretRepository
}

var (
	apiSecretKeyProviderOnce     sync.Once
	apiSecretKeyProviderInstance security.ApiSecretKeyProvider[*request.ApiContext]
)

func GetSecretKeyProvider() security.ApiSecretKeyProvider[*request.ApiContext] {
	apiSecretKeyProviderOnce.Do(func() {
		apiSecretKeyProviderInstance = &apiSecretKeyProviderImpl{
			repository: api_secret.GetRepository(),
		}
	})

	return apiSecretKeyProviderInstance
}

func (s *apiSecretKeyProviderImpl) Get(ctx *api_context.ApiRequestContext[*request.ApiContext]) (string, error) {
	apiSecret, err := s.repository.GetById(ctx.ApiKeyId)
	if err != nil {
		return "", err
	}
	return apiSecret.Key, nil
}
