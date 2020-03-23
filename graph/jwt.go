package graph

import (
	"context"
	"fmt"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

func getJwtClaimsCtx(ctx context.Context) (model.JwtClaims, error) {
	v := ctx.Value(model.JwtClaimsCtxKey)
	c, ok := v.(model.JwtClaims)
	if !ok {
		return model.JwtClaims{}, fmt.Errorf("Failed to get jwt Claims from context")
	}

	return c, nil
}

func getJwtTokenCtx(ctx context.Context) (string, error) {
	v := ctx.Value(model.JwtTokenKey)
	t, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("Failed to get jwt token from context")
	}

	return t, nil
}
