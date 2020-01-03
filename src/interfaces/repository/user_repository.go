package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/metadata"
)

type userRepository struct {
	client auth_grpc.AuthServiceClient
}

func NewUserRepository(c auth_grpc.AuthServiceClient) repository.UserRepository {
	return &userRepository{client: c}
}

func (r *userRepository) GetByID(ctx context.Context, id *auth_grpc.ID) (*auth_grpc.User, error) {
	return r.client.GetByID(ctx, id)
}

func (r *userRepository) Create(ctx context.Context, req *auth_grpc.CreateReq) (*auth_grpc.UserWithToken, error) {
	return r.client.Create(ctx, req)
}

func (r *userRepository) Update(ctx context.Context, req *auth_grpc.UpdateReq, token string) (*auth_grpc.User, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.client.Update(ctx, req)
}

func (r *userRepository) Delete(ctx context.Context, token string) (*wrappers.BoolValue, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.client.Delete(ctx, &empty.Empty{})
}

func (r *userRepository) Login(ctx context.Context, req *auth_grpc.LoginReq) (*auth_grpc.UserWithToken, error) {
	return r.client.Login(ctx, req)
}

func (r *userRepository) Logout(ctx context.Context, token string) (*wrappers.BoolValue, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.client.Logout(ctx, &empty.Empty{})
}

func (r *userRepository) RefreshIDToken(ctx context.Context, token string) (*auth_grpc.TokenPair, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return r.client.RefreshIDToken(ctx, &empty.Empty{})
}
