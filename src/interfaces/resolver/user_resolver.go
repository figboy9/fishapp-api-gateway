package resolver

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func (r *queryResolver) User(ctx context.Context, id string) (*graphql.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.UserInteractor.User(ctx, &auth_grpc.ID{Id: intID})
}

func (r *mutationResolver) CreateUser(ctx context.Context, in gen.CreateUserInput) (*gen.UserWithToken, error) {
	req := &auth_grpc.CreateReq{
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.CreateUser(ctx, req)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, in gen.UpdateUserInput) (*graphql.User, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	req := &auth_grpc.UpdateReq{
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.UpdateUser(ctx, req, token)
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return false, err
	}
	return r.UserInteractor.DeleteUser(ctx, token)
}

func (r *mutationResolver) Login(ctx context.Context, in gen.LoginInput) (*gen.UserWithToken, error) {
	req := &auth_grpc.LoginReq{
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.Login(ctx, req)
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return false, err
	}
	return r.UserInteractor.Logout(ctx, token)
}
func (r *mutationResolver) RefreshIDToken(ctx context.Context) (*graphql.TokenPair, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.UserInteractor.RefreshIDToken(ctx, token)
}
