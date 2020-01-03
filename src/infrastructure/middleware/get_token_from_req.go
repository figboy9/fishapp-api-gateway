package middleware

import (
	"context"
	"net/http"
	"strings"
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
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
