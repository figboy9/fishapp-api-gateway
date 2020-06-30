package middleware

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

func GetTokenFromWebsocketInit(ctx context.Context, p transport.InitPayload) (context.Context, error) {
	authHeader := p.Authorization()
	if authHeader == "" {
		return ctx, nil
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return context.WithValue(ctx, model.JwtTokenKey, token), nil
}
