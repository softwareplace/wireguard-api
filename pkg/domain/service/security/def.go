package security

import (
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"time"
)

type ApiSecurityService interface {
	Secret() []byte
	GenerateApiSecretJWT(jwtInfo ApiJWTInfo) (string, error)
	ExtractJWTClaims(requestContext server.ApiRequestContext) bool
	JWTClaims(ctx server.ApiRequestContext) (map[string]interface{}, error)
	GenerateJWT(user models.User) (map[string]interface{}, error)
	Encrypt(key string) (string, error)
	Decrypt(encrypted string) (string, error)
	Validation(
		ctx server.ApiRequestContext,
		next func(ctx server.ApiRequestContext) (*models.User, bool),
	) (*models.User, bool)
}

type ApiJWTInfo struct {
	Client string
	Key    string
	// Expiration in hours
	Expiration time.Duration //
}

type apiSecurityServiceImpl struct{}

var (
	instance = &apiSecurityServiceImpl{}
)

func GetApiSecurityService() ApiSecurityService {
	instance.Secret()
	return instance
}
