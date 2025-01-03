package security

import (
	"github.com/eliasmeireles/wireguard-api/pkg/handlers/request"
	"github.com/eliasmeireles/wireguard-api/pkg/models"
	"time"
)

type ApiSecurityService interface {
	Secret() []byte
	GenerateApiSecretJWT(jwtInfo ApiJWTInfo) (string, error)
	ExtractJWTClaims(requestContext *request.ApiRequestContext) bool
	JWTClaims(ctx *request.ApiRequestContext) (map[string]interface{}, error)
	GenerateJWT(user models.User) (map[string]string, error)
	Encrypt(key string) (string, error)
	Decrypt(encrypted string) (string, error)
	Validation(
		ctx *request.ApiRequestContext,
		next func(ctx *request.ApiRequestContext) (*models.User, bool),
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
