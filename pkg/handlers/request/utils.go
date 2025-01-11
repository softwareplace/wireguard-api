package request

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/pkg/models"
)

func (ctx *ApiRequestContext) GetRoles() (roles []string, err error) {
	user := ctx.GetAccessContext().User
	if user != nil && len(user.Roles) > 0 {
		return user.Roles, nil
	}
	return nil, fmt.Errorf("user roles not found")
}

func (ctx *ApiRequestContext) GetAccessContext() *AccessContext {
	accessUserContext := ctx.AccessContext

	if accessUserContext == nil {
		accessUserContext = &AccessContext{
			ApiKey:        ctx.Request.Header.Get(XApiKey),
			Authorization: ctx.Request.Header.Get("Authorization"),
		}
	}

	ctx.AccessContext = accessUserContext
	return accessUserContext
}

func (ctx *ApiRequestContext) SetUser(user *models.User) {
	ctx.GetAccessContext().User = user
}

func (ctx *ApiRequestContext) SetAuthorizationClaims(authorizationClaims map[string]interface{}) {
	ctx.GetAccessContext().AuthorizationClaims = authorizationClaims
}

func (ctx *ApiRequestContext) SetApiKeyClaims(apiKeyClaims map[string]interface{}) {
	ctx.GetAccessContext().ApiKeyClaims = apiKeyClaims
}

func (ctx *ApiRequestContext) SetApiKeyId(apiKeyId string) {
	ctx.GetAccessContext().ApiKeyId = apiKeyId
}

func (ctx *ApiRequestContext) SetAccessId(accessId string) {
	ctx.GetAccessContext().AccessId = accessId
}
