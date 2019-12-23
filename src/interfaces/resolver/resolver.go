package resolver

import (
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

type Resolver struct {
	PostInteractor interactor.UPostInteractor
}

func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

type contextKey string

const UserIDCtxKey contextKey = "userID"
