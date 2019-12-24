package resolver

import (
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

type Resolver struct {
	PostInteractor interactor.UPostInteractor
	UserInteractor interactor.UUserInteractor
}

func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Post() graphql.PostResolver {
	return &postResolver{r}
}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

type postResolver struct{ *Resolver }

type contextKey string

const UserIDCtxKey contextKey = "userID"
