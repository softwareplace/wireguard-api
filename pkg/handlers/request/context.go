package request

import (
	"context"
	"github.com/eliasmeireles/wireguard-api/pkg/models"
	"net/http"
)

const (
	AccessContextKey = "userAccessContext"
	xApiKey          = "X-Api-Key"
)

type ApiRequestContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type AccessContext struct {
	User                *models.User
	AccessId            string
	Authorization       string
	ApiKey              string
	ApiKeyId            string
	AuthorizationClaims map[string]interface{}
	ApiKeyClaims        map[string]interface{}
}

func Build(w http.ResponseWriter, r *http.Request) ApiRequestContext {
	ctx := ApiRequestContext{
		Writer:  w,
		Request: r,
	}

	ctx.GetAccessContext()
	return ctx
}

func (ctx *ApiRequestContext) GetAccessContext() *AccessContext {
	accessUserContext := ctx.Request.Context().Value(AccessContextKey)

	if accessUserContext == nil {
		accessUserContext = &AccessContext{
			ApiKey:        ctx.Request.Header.Get(xApiKey),
			Authorization: ctx.Request.Header.Get("Authorization"),
		}
	}
	apiRequestContext := context.WithValue(ctx.Request.Context(), AccessContextKey, accessUserContext)
	ctx.Request = ctx.Request.WithContext(apiRequestContext)

	return accessUserContext.(*AccessContext)
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
