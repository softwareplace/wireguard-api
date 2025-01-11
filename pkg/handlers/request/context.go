package request

import (
	"context"
	"github.com/google/uuid"
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

func Of(w http.ResponseWriter, r *http.Request) ApiRequestContext {
	currentContext := r.Context().Value(apiAccessContextKey)

	if currentContext != nil {
		return currentContext.(ApiRequestContext)
	}

	w.Header().Set("Content-Type", "application/json")
	ctx := ApiRequestContext{
		Writer:    w,
		Request:   r,
		sessionId: uuid.New().String(),
	}

	ctx.AccessContext = ctx.GetAccessContext()
	apiRequestContext := context.WithValue(ctx.Request.Context(), apiAccessContextKey, ctx)
	ctx.Request = r.WithContext(apiRequestContext)

	return ctx
}

func (ctx *ApiRequestContext) GetSessionId() string {
	return ctx.sessionId
}

func (ctx *ApiRequestContext) Next(next http.Handler) {
	next.ServeHTTP(ctx.Writer, ctx.Request)
}
