package request

import (
	"fmt"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"net/http"
)

type ApiContext struct {
	User                *models.User
	AccessId            string
	ApiKeyId            string
	AuthorizationClaims map[string]interface{}
	ApiKeyClaims        map[string]interface{}
}

func ContextBuilder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := server.Of(w, r, "MIDDLEWARE/CONTEXT_BUILDER")
		ctx.RequestData = &ApiContext{}
		ctx.Next(next)
	})
}

func (ctx *ApiContext) GetAccessApiKeyId() string {
	if ctx.ApiKeyId != "" {
		return ctx.ApiKeyId
	}
	return "N/A"
}

func (ctx *ApiContext) GetAccessId() string {
	if ctx.AccessId != "" {
		return ctx.AccessId
	}
	return "N/A"
}

func (ctx *ApiContext) new() *ApiContext {
	return &ApiContext{}
}

func (ctx *ApiContext) GetRoles() (roles []string, err error) {
	user := ctx.User
	if user != nil && len(user.Roles) > 0 {
		return user.Roles, nil
	}
	return nil, fmt.Errorf("user roles not found")
}

func (ctx *ApiContext) SetUser(user *models.User) {
	ctx.User = user
}

func (ctx *ApiContext) SetAuthorizationClaims(authorizationClaims map[string]interface{}) {
	ctx.AuthorizationClaims = authorizationClaims
}

func (ctx *ApiContext) SetApiKeyClaims(apiKeyClaims map[string]interface{}) {
	ctx.ApiKeyClaims = apiKeyClaims
}

func (ctx *ApiContext) SetApiKeyId(apiKeyId string) {
	ctx.ApiKeyId = apiKeyId
}

func (ctx *ApiContext) SetAccessId(accessId string) {
	ctx.AccessId = accessId
}
