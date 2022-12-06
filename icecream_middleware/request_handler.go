package icecream_middleware

import (
	"context"
	"github.com/rs/xid"
	"net/http"
)

func GenerateRequestIdHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = GenerateNewRequestId(r)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func GenerateNewRequestId(r *http.Request) *http.Request {
	requestId := xid.New()
	ctx := context.WithValue(r.Context(), "requestId", requestId.String())
	return r.WithContext(ctx)
}

func GetRequestIDFromRequest(r *http.Request) string {
	if r == nil {
		return "-"
	}
	return GetRequestIDFromContext(r.Context())
}

func GetRequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return "-"
	}
	value, ok := ctx.Value("requestId").(string)
	if ok {
		return value
	} else {
		return "-"
	}
}
