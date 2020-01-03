package resolver

import (
	"context"
	"fmt"

	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

type resolver struct {
	userInteractor interactor.UserInteractor
	postInteractor interactor.PostInteractor
}

func NewResolver(u interactor.UserInteractor, p interactor.PostInteractor) graphql.ResolverRoot {
	return &resolver{u, p}
}

func (r *resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

func (r *resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *resolver }

type mutationResolver struct{ *resolver }

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
