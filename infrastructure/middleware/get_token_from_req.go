package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler/transport"
)

type contextKey string

const jwtTokenKey contextKey = "jwtToken"

func (*middleware) GetTokenFromReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		ctx := context.WithValue(r.Context(), jwtTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (*middleware) GetTokenFromWebsocketInit(ctx context.Context, p transport.InitPayload) (context.Context, error) {
	authHeader := p.Authorization()
	if authHeader == "" {
		return ctx, nil
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return context.WithValue(ctx, jwtTokenKey, token), nil
}
