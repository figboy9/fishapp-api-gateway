package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const tokenCtxKey contextKey = "id-token"

func GetTokenFromReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		ctx := context.WithValue(r.Context(), tokenCtxKey, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
