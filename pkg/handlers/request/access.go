package request

import (
	"github.com/softwareplace/wireguard-api/pkg/models"
)

type AccessContext struct {
	User                *models.User
	AccessId            string
	Authorization       string
	ApiKey              string
	ApiKeyId            string
	AuthorizationClaims map[string]interface{}
	ApiKeyClaims        map[string]interface{}
}
