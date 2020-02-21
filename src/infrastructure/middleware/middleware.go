package middleware

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
)

type middleware struct{}

type Middleware interface {
	FieldMiddleware(ctx context.Context, next graphql.Resolver) (interface{}, error)
	GetTokenFromReq(next http.Handler) http.Handler
}

func InitMiddleware() Middleware {
	return &middleware{}
}
