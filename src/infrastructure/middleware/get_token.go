package middleware

import (
	"context"
	"net/http"
	"strings"
)

type reqToken struct{}

func GetToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]
		ctx := context.WithValue(r.Context(), reqToken{}, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
