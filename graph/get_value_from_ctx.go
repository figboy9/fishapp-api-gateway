package graph

import (
	"context"
	"errors"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

func getClaimsFromCtx(ctx context.Context) (model.JwtClaims, error) {
	v := ctx.Value(model.JwtClaimsCtxKey)
	c, ok := v.(model.JwtClaims)
	if !ok {
		return model.JwtClaims{}, errors.New("failed to get jwt claims from context")
	}

	return c, nil
}

func getTokenFromCtx(ctx context.Context) (string, error) {
	v := ctx.Value(model.JwtTokenKey)
	t, ok := v.(string)
	if !ok {
		return "", errors.New("failed to get jwt token from context")
	}

	return t, nil
}
