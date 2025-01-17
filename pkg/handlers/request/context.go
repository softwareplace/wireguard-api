package request

import (
	"fmt"
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/wireguard-api/pkg/models"
)

type ApiContext struct {
	User                *models.User
	AccessId            string
	ApiKeyId            string
	AuthorizationClaims map[string]interface{}
	ApiKeyClaims        map[string]interface{}
}

func (ctx *ApiContext) Data(data api_context.ApiContextData) {
	apiCtx := data.(*ApiContext)
	ctx.User = apiCtx.User
}

func (ctx *ApiContext) Salt() string {
	if ctx.User != nil {
		return ctx.User.Salt

	}
	return "N/A"
}

func (ctx *ApiContext) Roles() []string {
	if ctx.User != nil {
		return ctx.User.Roles
	}
	return []string{}
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
	return nil, fmt.Errorf("user_service roles not found")
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

func (ctx *ApiContext) SetRoles(roles []string) {
}

func ContextBuilder(ctx *api_context.ApiRequestContext[*ApiContext]) bool {
	ctx.RequestData = &ApiContext{}
	return true
}
