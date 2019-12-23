package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
)

type UserRepository interface {
	GetByID(ctx context.Context, id *user_grpc.ID) (*user_grpc.User, error)
}
