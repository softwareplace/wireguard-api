package request

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
)

const (
	apiAccessContextKey = "apiAccessContext"
	XApiKey             = "X-Api-Key"
)

type ApiRequestContext struct {
	Writer        http.ResponseWriter
	Request       *http.Request
	AccessContext *AccessContext
	sessionId     string
}

func Of(w http.ResponseWriter, r *http.Request, reference string) ApiRequestContext {
	currentContext := r.Context().Value(apiAccessContextKey)

	if currentContext != nil {
		var ctx *ApiRequestContext
		ctx = currentContext.(*ApiRequestContext)
		ctx.updateContext(r)
		return *ctx
	}

	return createNewContext(w, r, reference)
}

func createNewContext(w http.ResponseWriter, r *http.Request, reference string) ApiRequestContext {
	w.Header().Set("Content-Type", "application/json")
	ctx := ApiRequestContext{
		Writer:    w,
		Request:   r,
		sessionId: uuid.New().String(),
	}

	log.Printf("%s -> initialized a context with session id: %s", reference, ctx.sessionId)
	ctx.updateContext(r)
	return ctx
}

func (ctx *ApiRequestContext) updateContext(r *http.Request) {
	ctx.AccessContext = ctx.GetAccessContext()
	apiRequestContext := context.WithValue(ctx.Request.Context(), apiAccessContextKey, ctx)
	ctx.Request = r.WithContext(apiRequestContext)
}

func (ctx *ApiRequestContext) GetAccessApiKeyId() string {
	if ctx.AccessContext != nil && ctx.AccessContext.ApiKeyId != "" {
		return ctx.AccessContext.ApiKeyId
	}
	return "N/A"
}

func (ctx *ApiRequestContext) GetAccessId() string {
	if ctx.AccessContext != nil && ctx.AccessContext.AccessId != "" {
		return ctx.AccessContext.AccessId
	}
	return "N/A"
}

func (ctx *ApiRequestContext) GetSessionId() string {
	return ctx.sessionId
}

func (ctx *ApiRequestContext) Next(next http.Handler) {
	next.ServeHTTP(ctx.Writer, ctx.Request)
}
