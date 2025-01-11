package request

import (
	"context"
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
}

func Of(w http.ResponseWriter, r *http.Request) *ApiRequestContext {
	currentContext := r.Context().Value(apiAccessContextKey)

	if currentContext != nil {
		return currentContext.(*ApiRequestContext)
	}

	w.Header().Set("Content-Type", "application/json")
	ctx := ApiRequestContext{
		Writer:  w,
		Request: r,
	}

	ctx.AccessContext = ctx.GetAccessContext()
	apiRequestContext := context.WithValue(ctx.Request.Context(), ctx, apiAccessContextKey)
	ctx.Request = r.WithContext(apiRequestContext)

	return &ctx
}

func (ctx *ApiRequestContext) Next(next http.Handler) {
	next.ServeHTTP(ctx.Writer, ctx.Request)
}
