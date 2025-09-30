package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const requestIdHeaderName = "X-Request-Id"
const RequestIdContextAttributeName = "requestId"

func LoggerContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		xRequestId := r.Header.Get(requestIdHeaderName)
		if xRequestId == "" {
			xRequestId = uuid.NewString()
		}
		//lint:ignore SA1029 ignore
		ctx = context.WithValue(ctx, RequestIdContextAttributeName, xRequestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
