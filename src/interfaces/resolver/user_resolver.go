package resolver

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func (r *queryResolver) User(ctx context.Context, id string) (*graphql.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.UserInteractor.User(ctx, &user_grpc.ID{Id: intID})
}

func (r *mutationResolver) CreateUser(ctx context.Context, in gen.CreateUserInput) (*gen.UserWithToken, error) {
	createReq := &user_grpc.CreateReq{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.CreateUser(ctx, createReq)
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
