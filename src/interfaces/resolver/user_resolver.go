package resolver

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func (r *queryResolver) User(ctx context.Context, id string) (*graphql.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, in gen.CreateUserInput) (*gen.UserWithToken, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, in gen.UpdateUserInput) (*gen.UserWithToken, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, in gen.LoginInput) (*gen.UserWithToken, error) {
	panic("not implemented")
}
