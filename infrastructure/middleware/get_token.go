package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

func GetTokenFromHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		ctx := context.WithValue(r.Context(), model.JwtTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
