package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type UserRepository struct {
	Client user_grpc.UserServiceClient
}

func (r *UserRepository) GetByID(ctx context.Context, id *user_grpc.ID) (*user_grpc.User, error) {
	return r.Client.GetByID(ctx, id)
}

func (r *UserRepository) Create(ctx context.Context, req *user_grpc.CreateReq) (*user_grpc.UserWithToken, error) {
	return r.Client.Create(ctx, req)
}

func (r *UserRepository) Update(ctx context.Context, req *user_grpc.UpdateReq) (*user_grpc.User, error) {
	return r.Client.Update(ctx, req)
}

func (r *UserRepository) Delete(ctx context.Context, id *user_grpc.ID) (*wrappers.BoolValue, error) {
	return r.Client.Delete(ctx, id)
}

func (r *UserRepository) Login(ctx context.Context, req *user_grpc.LoginReq) (*user_grpc.UserWithToken, error) {
	return r.Client.Login(ctx, req)
}

func (r *UserRepository) AddBlackList(ctx context.Context, req *user_grpc.AddBlackListReq) (*wrappers.BoolValue, error) {
	return r.Client.AddBlackList(ctx, req)
}
func (r *UserRepository) CheckBlackListAndGenToken(ctx context.Context, req *user_grpc.CheckBlackListReq) (*user_grpc.TokenPair, error) {
	return r.Client.CheckBlackListAndGenToken(ctx, req)
}
