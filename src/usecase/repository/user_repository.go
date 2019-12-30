package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type UserRepository interface {
	GetByID(ctx context.Context, id *user_grpc.ID) (*user_grpc.User, error)
	Create(ctx context.Context, req *user_grpc.CreateReq) (*user_grpc.UserWithToken, error)
	Update(ctx context.Context, req *user_grpc.UpdateReq) (*user_grpc.User, error)
	Delete(ctx context.Context, req *user_grpc.ID) (*wrappers.BoolValue, error)
	Login(ctx context.Context, req *user_grpc.LoginReq) (*user_grpc.UserWithToken, error)
	AddBlackList(ctx context.Context, in *user_grpc.AddBlackListReq) (*wrappers.BoolValue, error)
	CheckBlackListAndGenToken(ctx context.Context, in *user_grpc.CheckBlackListReq) (*user_grpc.TokenPair, error)
}
