package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type UserRepository interface {
	GetByID(ctx context.Context, id *auth_grpc.ID) (*auth_grpc.User, error)
	Create(ctx context.Context, req *auth_grpc.CreateReq) (*auth_grpc.UserWithToken, error)
	Update(ctx context.Context, req *auth_grpc.UpdateReq, token string) (*auth_grpc.User, error)
	Delete(ctx context.Context, token string) (*wrappers.BoolValue, error)
	Login(ctx context.Context, req *auth_grpc.LoginReq) (*auth_grpc.UserWithToken, error)
	Logout(ctx context.Context, token string) (*wrappers.BoolValue, error)
	RefreshIDToken(ctx context.Context, token string) (*auth_grpc.TokenPair, error)
}
