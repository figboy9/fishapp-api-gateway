package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const idTokenCtxKey contextKey = "id-token"

func GetTokenFromReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]
		ctx := context.WithValue(r.Context(), idTokenCtxKey, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
