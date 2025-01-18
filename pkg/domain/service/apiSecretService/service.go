package apiSecretService

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/api_secret"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
)

type ApiSecretService interface {
	GetKey(ctx *api_context.ApiRequestContext[*request.ApiContext]) (string, error)
}

type serviceImpl struct {
	repository api_secret.ApiSecretRepository
}

func GetService() ApiSecretService {
	return &serviceImpl{
		repository: api_secret.GetRepository(),
	}
}

func (s *serviceImpl) GetKey(ctx *api_context.ApiRequestContext[*request.ApiContext]) (string, error) {
	apiSecret, err := s.repository.GetById(ctx.ApiKeyId)
	if err != nil {
		return "", err
	}
	return apiSecret.Key, nil
}
