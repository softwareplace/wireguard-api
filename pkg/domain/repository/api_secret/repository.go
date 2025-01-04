package api_secret

import (
	"github.com/softwareplace/wireguard-api/pkg/models"
)

type ApiSecretRepository interface {
	Save(apiSecret models.ApiSecret) (*string, error)
	GetById(id string) (models.ApiSecret, error)
}

type repositoryImpl struct{}

func GetRepository() ApiSecretRepository {
	return &repositoryImpl{}
}
