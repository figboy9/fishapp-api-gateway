package middleware

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

type middleware struct{}

type Middleware interface {
	AuthMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver, authAPI bool) (res interface{}, err error)
	GetTokenFromReq(next http.Handler) http.Handler
	GetTokenFromWebsocketInit(ctx context.Context, p transport.InitPayload) (context.Context, error)
}

func InitMiddleware() Middleware {
	return &middleware{}
}
