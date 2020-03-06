// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/auth_grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	res, err := r.authClient.Create(ctx, &auth_grpc.CreateReq{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	u, err := convertUserGQL(res.User)
	if err != nil {
		return nil, err
	}
	return &model.CreateUserPayload{
		User:      u,
		TokenPair: convertTokenPairGQL(res.TokenPair),
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (*model.DeleteUserPayload, error) {
	t, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.authClient.Delete(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&empty.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return &model.DeleteUserPayload{Success: res.Value}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.UpdateUserPayload, error) {
	t, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.authClient.Update(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&auth_grpc.UpdateReq{
			Email:    input.Email,
			Password: input.Password,
		})
	if err != nil {
		return nil, err
	}
	u, err := convertUserGQL(res)
	if err != nil {
		return nil, err
	}
	return &model.UpdateUserPayload{User: u}, nil
}

func (r *mutationResolver) RefreshIDToken(ctx context.Context) (*model.RefreshIDTokenPayload, error) {
	t, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.authClient.RefreshIDToken(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&empty.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return &model.RefreshIDTokenPayload{TokenPair: convertTokenPairGQL(res)}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	res, err := r.authClient.Login(ctx, &auth_grpc.LoginReq{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	u, err := convertUserGQL(res.User)
	if err != nil {
		return nil, err
	}
	return &model.LoginPayload{User: u, TokenPair: convertTokenPairGQL(res.TokenPair)}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*model.LogoutPayload, error) {
	t, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.authClient.Logout(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&empty.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return &model.LogoutPayload{Success: res.Value}, nil
}

func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	res, err := r.authClient.GetByID(ctx, &auth_grpc.ID{Id: id})
	if err != nil {
		return nil, err
	}
	u, err := convertUserGQL(res)
	if err != nil {
		return nil, err
	}
	return u, nil
}
