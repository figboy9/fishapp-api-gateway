package graph

import (
	"context"
	"fmt"
)

type contextKey string

type JwtClaims struct {
	UserID    int64
	Jti       string
	ExpiresAt int64
}

const JwtClaimsCtxKey contextKey = "jwtClaims"

const JwtTokenKey contextKey = "jwtToken"

func getJwtClaimsCtx(ctx context.Context) (JwtClaims, error) {
	v := ctx.Value(JwtClaimsCtxKey)
	jwtClaims, ok := v.(JwtClaims)
	if !ok {
		return JwtClaims{}, fmt.Errorf("Failed to get jwt Claims from context")
	}

	return jwtClaims, nil
}

func getJwtTokenCtx(ctx context.Context) (string, error) {
	v := ctx.Value(JwtTokenKey)
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("Failed to get jwt token from context")
	}

	return token, nil
}
