package resolver

import (
	"context"
	"fmt"

	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

type resolver struct {
	// sync.RWMutex
	// gqlChans          map[int64]chan *graphql.Message
	userInteractor    interactor.UserInteractor
	profileInteractor interactor.ProfileInteractor
	postInteractor    interactor.PostInteractor
	chatInteractor    interactor.ChatInteractor
}

func NewResolver(
	u interactor.UserInteractor,
	pri interactor.ProfileInteractor,
	pi interactor.PostInteractor,
	ci interactor.ChatInteractor,
) gen.ResolverRoot {
	return &resolver{
		// RWMutex:           sync.RWMutex{},
		// gqlChans:          map[int64]chan *graphql.Message{},
		userInteractor:    u,
		profileInteractor: pri,
		postInteractor:    pi,
		chatInteractor:    ci,
	}
}

func (r *resolver) EntryPost() gen.EntryPostResolver {
	return &entryPostResolver{r}
}

func (r *resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}
func (r *resolver) Post() gen.PostResolver {
	return &postResolver{r}
}
func (r *resolver) Profile() gen.ProfileResolver {
	return &profileResolver{r}
}
func (r *resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}
func (r *resolver) Subscription() gen.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type queryResolver struct{ *resolver }

type mutationResolver struct{ *resolver }

type profileResolver struct{ *resolver }

type postResolver struct{ *resolver }

type entryPostResolver struct{ *resolver }

type subscriptionResolver struct{ *resolver }

type chatRoomResolver struct{ *resolver }

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
