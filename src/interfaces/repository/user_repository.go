package repository

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
)

type UserRepository struct {
	Client user_grpc.UserServiceClient
}
