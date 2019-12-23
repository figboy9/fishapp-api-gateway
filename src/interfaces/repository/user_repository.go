package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
)

type UserRepository struct {
	Client user_grpc.UserServiceClient
}

func (r *UserRepository) GetByID(ctx context.Context, id *user_grpc.ID) (*user_grpc.User, error) {
	return r.Client.GetByID(ctx, id)
}
