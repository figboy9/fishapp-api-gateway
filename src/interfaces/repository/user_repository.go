package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/metadata"
)

type UserRepository struct {
	Client auth_grpc.AuthServiceClient
}

func (r *UserRepository) GetByID(ctx context.Context, id *auth_grpc.ID) (*auth_grpc.User, error) {
	return r.Client.GetByID(ctx, id)
}

func (r *UserRepository) Create(ctx context.Context, req *auth_grpc.CreateReq) (*auth_grpc.UserWithToken, error) {
	return r.Client.Create(ctx, req)
}

func (r *UserRepository) Update(ctx context.Context, req *auth_grpc.UpdateReq, token string) (*auth_grpc.User, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.Client.Update(ctx, req)
}

func (r *UserRepository) Delete(ctx context.Context, token string) (*wrappers.BoolValue, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.Client.Delete(ctx, &empty.Empty{})
}

func (r *UserRepository) Login(ctx context.Context, req *auth_grpc.LoginReq) (*auth_grpc.UserWithToken, error) {
	return r.Client.Login(ctx, req)
}

func (r *UserRepository) Logout(ctx context.Context, token string) (*wrappers.BoolValue, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.Client.Logout(ctx, &empty.Empty{})
}

func (r *UserRepository) RefreshIDToken(ctx context.Context, token string) (*auth_grpc.TokenPair, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.Client.RefreshIDToken(ctx, &empty.Empty{})
}
